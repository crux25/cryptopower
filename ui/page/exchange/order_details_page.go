package exchange

import (
	"context"
	"fmt"

	"gioui.org/layout"
	"gioui.org/unit"

	"code.cryptopower.dev/group/cryptopower/app"
	"code.cryptopower.dev/group/cryptopower/libwallet/instantswap"
	"code.cryptopower.dev/group/cryptopower/libwallet/utils"
	"code.cryptopower.dev/group/cryptopower/ui/cryptomaterial"
	"code.cryptopower.dev/group/cryptopower/ui/load"
	"code.cryptopower.dev/group/cryptopower/ui/page/components"
	"code.cryptopower.dev/group/cryptopower/ui/values"

	api "code.cryptopower.dev/exchange/instantswap"
)

const OrderDetailsPageID = "OrderDetailsPage"

type OrderDetailsPage struct {
	*load.Load
	// GenericPageModal defines methods such as ID() and OnAttachedToNavigator()
	// that helps this Page satisfy the app.Page interface. It also defines
	// helper methods for accessing the PageNavigator that displayed this page
	// and the root WindowNavigator.
	*app.GenericPageModal

	ctx       context.Context // page context
	ctxCancel context.CancelFunc

	exchange  api.IDExchange
	orderInfo *instantswap.Order

	backButton     cryptomaterial.IconButton
	infoButton     cryptomaterial.IconButton
	refreshBtn     cryptomaterial.Button
	createOrderBtn cryptomaterial.Button
}

func NewOrderDetailsPage(l *load.Load, order *instantswap.Order) *OrderDetailsPage {
	pg := &OrderDetailsPage{
		Load:             l,
		GenericPageModal: app.NewGenericPageModal(OrderDetailsPageID),
		orderInfo:        order,
	}

	exchange, err := pg.WL.MultiWallet.InstantSwap.NewExchanageServer(order.Server)
	if err != nil {
		fmt.Println(err)
	}
	pg.exchange = exchange

	pg.backButton, _ = components.SubpageHeaderButtons(l)

	_, pg.infoButton = components.SubpageHeaderButtons(pg.Load)

	pg.createOrderBtn = pg.Theme.Button(values.String(values.StrCreateNewOrder))
	pg.refreshBtn = pg.Theme.Button(values.String(values.StrRefresh))

	pg.orderInfo, err = pg.getOrderInfo(pg.orderInfo.UUID)

	return pg
}

func (pg *OrderDetailsPage) ID() string {
	return OrderDetailsPageID
}

func (pg *OrderDetailsPage) OnNavigatedTo() {
	pg.ctx, pg.ctxCancel = context.WithCancel(context.TODO())
}

func (pg *OrderDetailsPage) OnNavigatedFrom() {
	if pg.ctxCancel != nil {
		pg.ctxCancel()
	}
}

func (pg *OrderDetailsPage) HandleUserInteractions() {
	if pg.refreshBtn.Clicked() {
		pg.orderInfo, _ = pg.getOrderInfo(pg.orderInfo.UUID)
	}
}

func (pg *OrderDetailsPage) Layout(gtx C) D {
	container := func(gtx C) D {
		sp := components.SubPage{
			Load:       pg.Load,
			Title:      values.String(values.StrOrderDetails),
			BackButton: pg.backButton,
			Back: func() {
				pg.ParentNavigator().CloseCurrentPage()
			},
			Body: pg.layout,
		}
		return sp.Layout(pg.ParentWindow(), gtx)
	}

	return components.UniformPadding(gtx, container)
}

func (pg *OrderDetailsPage) layout(gtx C) D {
	return cryptomaterial.LinearLayout{
		Width:     cryptomaterial.MatchParent,
		Height:    cryptomaterial.MatchParent,
		Direction: layout.Center,
	}.Layout2(gtx, func(gtx C) D {
		return cryptomaterial.LinearLayout{
			Width:     gtx.Dp(values.MarginPadding550),
			Height:    cryptomaterial.MatchParent,
			Direction: layout.W,
			Margin: layout.Inset{
				Bottom: values.MarginPadding30,
			},
		}.Layout2(gtx, func(gtx C) D {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx C) D {
					return layout.Inset{
						Bottom: values.MarginPadding16,
					}.Layout(gtx, func(gtx C) D {
						return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
							layout.Rigid(func(gtx C) D {
								return layout.E.Layout(gtx, func(gtx C) D {
									return layout.Flex{
										Axis:      layout.Horizontal,
										Alignment: layout.Middle,
									}.Layout(gtx,
										layout.Rigid(func(gtx C) D {
											return layout.Inset{Right: values.MarginPadding2}.Layout(gtx, func(gtx C) D {
												return pg.Theme.Card().Layout(gtx, func(gtx C) D {
													return layout.UniformInset(values.MarginPadding20).Layout(gtx, func(gtx C) D {
														return layout.Center.Layout(gtx, func(gtx C) D {
															return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
																layout.Rigid(func(gtx C) D {
																	return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
																		layout.Rigid(func(gtx C) D {
																			return components.SetWalletLogo(pg.Load, gtx, pg.orderInfo.FromCurrency, values.MarginPadding30)
																		}),
																		layout.Rigid(func(gtx C) D {
																			return layout.Inset{
																				Left: values.MarginPadding10,
																			}.Layout(gtx, func(gtx C) D {
																				return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
																					layout.Rigid(func(gtx C) D {
																						return pg.Theme.Label(values.TextSize16, values.String(values.StrSending)).Layout(gtx)
																					}),
																					layout.Rigid(func(gtx C) D {
																						return components.LayoutOrderAmount(pg.Load, gtx, pg.orderInfo.FromCurrency, pg.orderInfo.InvoicedAmount)
																					}),
																					layout.Rigid(func(gtx C) D {
																						sourceWallet := pg.WL.MultiWallet.WalletWithID(pg.orderInfo.SourceWalletID)
																						sourceWalletName := sourceWallet.GetWalletName()
																						sourceAccount, _ := sourceWallet.GetAccount(pg.orderInfo.SourceAccountNumber)
																						fromText := fmt.Sprintf(values.String(values.StrOrderSendingFrom), sourceWalletName, sourceAccount.Name)
																						return pg.Theme.Label(values.TextSize16, fromText).Layout(gtx)
																					}),
																				)
																			})
																		}),
																	)
																}),
																layout.Rigid(func(gtx C) D {
																	return layout.Inset{
																		Top:    values.MarginPadding24,
																		Bottom: values.MarginPadding24,
																	}.Layout(gtx, func(gtx C) D {
																		return pg.Theme.Icons.ArrowDownIcon.LayoutSize(gtx, values.MarginPadding60)
																	})
																}),
																layout.Rigid(func(gtx C) D {
																	return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
																		layout.Rigid(func(gtx C) D {
																			return components.SetWalletLogo(pg.Load, gtx, pg.orderInfo.ToCurrency, values.MarginPadding30)
																		}),
																		layout.Rigid(func(gtx C) D {
																			return layout.Inset{
																				Left: values.MarginPadding10,
																			}.Layout(gtx, func(gtx C) D {
																				return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
																					layout.Rigid(func(gtx C) D {
																						return pg.Theme.Label(values.TextSize16, values.String(values.StrReceiving)).Layout(gtx)
																					}),
																					layout.Rigid(func(gtx C) D {
																						return components.LayoutOrderAmount(pg.Load, gtx, pg.orderInfo.ToCurrency, pg.orderInfo.OrderedAmount)
																					}),
																					layout.Rigid(func(gtx C) D {
																						destinationWallet := pg.WL.MultiWallet.WalletWithID(pg.orderInfo.DestinationWalletID)
																						destinationWalletName := destinationWallet.GetWalletName()
																						destinationAccount, _ := destinationWallet.GetAccount(pg.orderInfo.DestinationAccountNumber)
																						toText := fmt.Sprintf(values.String(values.StrOrderReceivingTo), destinationWalletName, destinationAccount.Name)
																						return pg.Theme.Label(values.TextSize16, toText).Layout(gtx)
																					}),
																					layout.Rigid(func(gtx C) D {
																						return pg.Theme.Label(values.TextSize16, pg.orderInfo.DestinationAddress).Layout(gtx)
																					}),
																				)
																			})
																		}),
																	)
																}),
															)
														})

													})
												})
											})
										}),
									)
								})
							}),
						)
					})
				}),
				layout.Rigid(func(gtx C) D {
					return pg.Theme.Label(values.TextSize28, pg.orderInfo.Status.String()).Layout(gtx)
				}),
				layout.Rigid(func(gtx C) D {
					if pg.orderInfo.Status == api.OrderStatusWaitingForDeposit {
						return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
							layout.Rigid(func(gtx C) D {
								return pg.Theme.Label(values.TextSize18, values.String(values.StrExpiresIn)).Layout(gtx)
							}),
							layout.Rigid(func(gtx C) D {
								return pg.Theme.Label(values.TextSize18, fmt.Sprint(pg.orderInfo.ExpiryTime)).Layout(gtx)
							}),
						)
					}
					return D{}
				}),
				layout.Rigid(func(gtx C) D {
					return layout.E.Layout(gtx, func(gtx C) D {
						return layout.Inset{
							Top: values.MarginPadding16,
						}.Layout(gtx, func(gtx C) D {
							return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
								layout.Rigid(func(gtx C) D {
									return pg.refreshBtn.Layout(gtx)
								}),
								layout.Rigid(func(gtx C) D {
									return layout.Inset{
										Left: values.MarginPadding10,
									}.Layout(gtx, func(gtx C) D {
										return pg.createOrderBtn.Layout(gtx)
									})
								}),
							)
						})
					})
				}),
			)
		})
	})
}

func (pg *OrderDetailsPage) getOrderInfo(UUID string) (*instantswap.Order, error) {
	orderInfo, err := pg.WL.MultiWallet.InstantSwap.GetOrderInfo(pg.exchange, UUID)
	if err != nil {
		return nil, err
	}

	return orderInfo, nil
}

func (pg *OrderDetailsPage) setWalletLogo(gtx C, currency string, size unit.Dp) D {
	if currency == utils.DCRWalletAsset.String() {
		return pg.Theme.Icons.DecredSymbol2.LayoutSize(gtx, values.MarginPadding40)
	}
	return pg.Theme.Icons.BTC.LayoutSize(gtx, values.MarginPadding40)
}
