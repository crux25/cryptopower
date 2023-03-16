package localizable

const ENGLISH = "en"

// one string per line, no multiline
// semicolon is not compulsory
const EN = `
"abandoned" = "Abandoned"
"about" = "About"
"abstain" = "Abstain"
"account" = "Account"
"accountList" = "Account List"
"accountMixer" = "AccountMixer"
"accRenamed" = "Account renamed"
"acctCreated" = "Account created"
"acctCreated" = "Account created"
"acctDetailsKey" = "%d external, %d internal, %d imported"
"acctName" = "Account name"
"acctNum" = "Account Number"
"acctRenamed" = "Account renamed"
"addAcctWarn" = "%v Accounts %v cannot %v be deleted once created.%v"
"addDexServer" = "Add dex server"
"addNewAccount" = "Add account"
"address" = "Address"
"addressCopied" = "Address copied"
"addressDiscoveryInProgress" = "Address Discovery in Progress..."
"addressDiscoveryStarted" = "Address discovery started successfully"
"addressDiscoveryStartedBody"    = "See wallet information page for progress"
"addrNotOwned" = "Address not owned by any wallet"
"addVSP" = "Add a new VSP..."
"addWallet" = "Add wallet"
"adminToTriggerVoting" = "Waiting for admin to trigger the start of voting"
"agendas" = "Agendas"
"ago" = "ago"
"all" = "All"
"allowSpendingFromUnmixedAccount" = "Allow spending from unmixed account"
"allowUnspendUnmixedAcct" = "%v Spendings from unmixed accounts could potentially be traced back to you %v Please type %v I am aware of the risks %v to allow spending from unmixed accounts.%v"
"allTickets" = "All tickets"
"amount" = "Amount"
"appLog" = "Application log"
"appName" = "Cryptopower"
"approved" = "Approved"
"appTitle" = "Cryptopower (%s)"
"appWallet" = "Cryptopower Wallet"
"askedEnterSeedWords" = "You will be asked to enter the seed phrase on the next screen."
"authorToAuthorizeVoting" = "Waiting for author to authorize voting"
"automatic" = "Automatic"
"autoSetUp" = "Auto Setup"
"autoSyncInfo" = "Auto sync feature has been enable, and wallets are not synced.\nWould you like to start syncing your wallets now?"
"autoTicketInfo" = "Cryptopower must remain running, for tickets to be automatically purchased"
"autoTicketPurchase" = "Auto ticket purchase"
"autoTicketWarn" = "Settings can not be modified when ticket buyer is running."
"backAndRename" = "Go back & rename"
"backStaking" = "Back to staking"
"backToWallets" = "Back to Wallets"
"backupInfo" = "%v No backup - no coins! %v In order not to lose your coins when your device is lost or broken, please make a wallet backup %v Now %v and keep it in %v a safe place! %v"
"backupLater" = "Backup later"
"backupNow" = "Backup now"
"backupSeedPhrase" = "Back up seed phrase"
"backupWarning" = "Wallet backup needed"
"balance" = "Balance:"
"balanceAfter" = "Balance after send"
"balToMaintain" = "Balance to maintain (DCR)"
"balToMaintainValue" = "Balance to maintain: %2.f"
"beepForNewBlocks" = "Beep for new blocks"
"bestBlockAge" = "Best block age"
"bestBlocks" = "Best block"
"bestBlockTimestamp" = "Best block timestamp"
"blockHeaderFetched" = "Block header fetched"
"blockHeaderFetchedCount" = "%d of %d"
"blocksLeft" = "%d blocks left"
"blocksScanned" = "Blocks scanned"
"build" = "Build"
"buildDate" = "Build date"
"canBuy" = "Can Buy"
"cancel" = "Cancel"
"canceling" = "Cancelling..."
"cancelMixer" = "Cancel mixer?"
"change" = "Change"
"changeAccount" = "Change account"
"changeSpecificPeer" = "Change specific peer"
"changeSpendingPass" = "Change spending passphrase"
"changeStartupPassword" = "Change startup password"
"changeUserAgent" = "Change user agent"
"changeWalletName" = "Change wallet name"
"checkGovernace" = "Check Governance page"
"checkMixerStatus" = "Check mixer status"
"checkStatistics" = "Check statistics"
"checkWalletLog" = "Check wallet logs"
"clear" = "Clear"
"clearAll" = "Clear all"
"clearSelection" = "Clear Selection"
"closingWallet" = "Shutting down Wallets..."
"coinSelection" = "Coin Selection"
"colon" = ": "
"complete" = "Completed"
"confirm" = "Confirm"
"confirmations" = "Confirmations"
"confirmDexReset" = "Confirm DEX Client Reset"
"confirmed" = "Confirmed"
"confirmNewSpendingPassword" = "Confirm new spending passphrase"
"confirmNewStartupPass" = "Confirm new startup password"
"confirmOrder" = "Confirm Order"
"confirmPending" = "Confirmation Required"
"confirmPurchase" = "Confirm Automatic Ticket Purchase"
"confirmRemoveStartupPass" = "Confirm to turn off startup password"
"confirmSend" = "Confim to send"
"confirmSpendingPassword" = "Confirm spending passphrase"
"confirmStartupPass" = "Confirm current startup password"
"confirmtoCreateAccs" = "Confirm to create needed accounts"
"confirmToMixAcc" = "Confirm to mix account"
"confirmToRemove" = "Confirm to remove"
"confirmToSetMixer" = "Confirm to set mixer accounts"
"confirmToShowSeed" = "Confirm to show seed"
"confirmToSign" = "Confirm to sign"
"confirmToVerifySeed" = "Confirm to verify seed"
"confirmUmixedSpending" = "Confirm to allow spending from unmixed accounts"
"confirmVote" = "Confirm your vote"
"confirmYourOrder" = "Confirm your order"
"confStatus" = "Confirmation Status"
"connectedPeersCount" = "Connected peers count"
"connectedTo" = "connected to %v peers"
"connecting" = "Connecting..."
"connection" = "Connection"
"connectToSpecificPeer" = "Connect to specific peer"
"consensusChange" = "Consensus Changes"
"consensusDashboard" = "Consensus Vote Dashboard"
"continue" = "Continue"
"coordinationServer" = "Coordination server"
"copied" = "Copied!"
"copy" = "Copy"
"copyBlockLink" = "Copy block explorer link"
"copyLink" = "Copy and paste the link below in your browser."
"copyseed" = "Copy seed"
"cost" = "Cost%v"
"create" = "Create"
"createANewWallet" = "Create a new wallet"
"createNewAccount" = "Create new account"
"createNewOrder" = "Create new order"
"createNSetUpAccs" = "Create and setup the needed accounts for you."
"createOrder" = "Create Order"
"createOrderPageInfo" = "To change the default source and destination wallet/account used for exchange, click the settings icon."
"createStartupPassword" = "Create a startup password"
"createWallet" = "Create wallet"
"currentSpendingPassword" = "Current spending passphrase"
"currentStartupPass" = "Current startup password"
"currentTotalBalance" = "Current Total Balance"
"CustomUserAgent" = "Custom user agent"
"dangerZone" = "Danger zone"
"darkMode" = "Dark mode"
"dateCreated" = "Date Created"
"dateSize" = "Wallet data"
"dayAgo" = "%d day ago"
"days" = "Days"
"daysAgo" = "%d days ago"
"daysToMiss" = "Days to miss"
"daysToVote" = "Days to vote"
"dcrCaps"            = "DCR"
"dcrReceived" = "You have received %s DCR"
"debug" = "Debug"
"default" = "default"
"delete" = "Delete"
"descriptionNote" = "Description Note"
"destAddr" = "Destination Address"
"destination" = "Destination"
"destinationMissing" = "destination missing"
"destinationModalInfo" = "A new receiving address will be automatically generated within the selected account."
"destinationWalletNotSynced" = "Destination wallet is not synced"
"dex" = "Dex"
"dexDataReset" = "DEX client data reset complete."
"dexDataResetFalse" = "DEX client data reset failed. Check the logs."
"dexResetInfo" = "You may need to restart cryptopower before you can use the DEX again. Proceed?"
"dexStartupErr" = "Unable to start DEX client: %v"
"disable" = "Disable"
"disabled" = "disabled"
"disconnect" = "Disconnect"
"discoverAddressUsage" = "Discover Address Usage"
"discoveringWalletAddress" = "Discovering wallet address · %v%%"
"discussions" = "Discussions:   %d comments"
"documentation" = "Documentation"
"done" = "Done"
"duration" = "%s (%d/%d blocks)"
"edit" = "Edit"
"emptyMsg" = "Field cannot be empty. Please provide valid signed message."
"emptySign" = "Field cannot be empty. Please provide valid signature."
"enabled" = "enabled"
"english" = "English"
"enterAddressToSign" = "Enter an address and message to sign:"
"enterHex"       = "Enter Hex"
"enterSeedPhrase" = "Enter your seed phrase"
"enterSpendingPassword" = "Enter spending passphrase"
"enterValidAddress" = "Please enter a valid address"
"enterValidMsg" = "Please enter a valid message to sign"
"enterWalletDetails" = "Enter wallet details"
"enterWalletName" = "Enter wallet name"
"enterWalletSeed" = "Enter wallet seed"
"enterXpubKey" = "Enter a valid extended PubKey"
"errNoMixable" = "No mixable output"
"errorMovingFunds" = "Error moving funds: Auto funds transfer has been skipped. Move funds to unmixed account manually from the send page. \n%s"
"errPassEmpty" = "Password can not be empty"
"estimatedSize" = "Estimated size"
"estimatedTime" = "Estimated time"
"exchange" = "Exchange"
"exchangeAPI" = "Exchange API"
"exchangeIntro" = "Exchange curriencies simple, fast and secure."
"exchangeRate" = "Fetch Exchange Rate"
"existingWalletName" = "What is your wallet existing wallet name?"
"exit" = "Exit"
"expired" = "Expired"
"expiredInfo" = "This Stake has not been chosen to vote within %d blocks, and thus expired"
"expiredInfoDisc" = "Expired tickets will be revoked to return the original Stake price to you"
"expiredInfoDiscSub" = "If a Stake is not revoked automatically, use the revoke button."
"expiredOn" = "Expired on"
"expiresIn" = "Expires in "
"explorerURL" = "Explorer URL for %v Asset"
"extendedInfo" = "The Extended Public Key is used to import the wallet as a watch-only wallet"
"extendedKey" = "Extended Public Key"
"extendedKeyCopied" = "Extended Public Key copied"
"extendedPubKey" = "Extended public key"
"external" = "External"
"failed" = "Failed"
"fee" = "Fee"
"feeRateAPI" = "Fee Rates API"
"feerates" = "Fee Rates"
"fetchingAgenda" = "Fetching agendas..."
"fetchingBlockHeaders" = "Fetching block headers · %v%%"
"fetchingOrders" = "Fetching Orders"
"fetchingPolicies" = "Fetching policies"
"fetchingProposals" = "Fetching proposals..."
"fetchProposals" = "Fetch proposals"
"fetchRateError" = "error fetching rate"
"fetchRates" = "Fetch Rates"
"finished" = "Finished"
"french" = "French"
"from" = "From"
"functionUnavailable" = "This function is unavailable until sync is complete."
"gapLimit" = "Gap Limit"
"gapLimitInputErr" = "Invalid input: valid values (1-1000)"
"general" = "General"
"generateAddress" = "Generate new address"
"gotIt" = "Got it"
"governance" = "Governance"
"governanceAPI" = "Governance API(s)"
"governanceInfo" = "%vThe Decred community can participate in proposal discussions for new initiatives and request funding for these initiatives. Decred stakeholders can vote if these proposals should be approved and paid for by the Decred Treasury. %v Would you like to fetch and view the proposals?%v"
"governanceSettingsInfo" = "Are you sure you want to disable governance? This will clear all available proposals"
"hash" = "Hash"
"hdPath" = "HD Path"
"help" = "Help"
"helpInfo" = "For more information, click and please visit the Decred documentation."
"hex" = "Hex"
"hideDetails" = "Hide details"
"hideSeedPhrase" = "Anyone with your seed phrase can steal your funds. DO NOT show it to anyone."
"hint" = "Hint"
"history" = "History"
"hourAgo" = "%d hour ago"
"hours" = "Hours"
"hoursAgo" = " %d hours ago"
"howGovernanceWork" = "How does Governance Work?"
"howNotToStoreSeedPhrase" = "It is highly discouraged to store your seed phrase in any digital format (e.g. screenshot)."
"howToCopy" = "How to copy"
"howToStoreSeedPhrase" = "It is recommended to store your seed phrase in a physical format (e.g. write down on a paper)."
"httpReq" = "For HTTP request"
"imawareOfRisk" = "I am aware of the risks"
"immature" = "Immature"
"immatureInfo" = "Mature in %v of %v blocks (%v)"
"immatureRewards" = "Immature Rewards"
"immatureStakeGen" = "Immature Stake Gen"
"import" = "Import"
"importantSeedPhrase" = "The 33-word seed phrase is EXTREMELY IMPORTANT."
"imported" = "imported"
"importExistingWallet" = "Import an existing wallet"
"importWatchingOnlyWallet" = "Import a watch-only wallet"
"includedInBlock" = "Included in block"
"inDiscussion" = "In discussion"
"info" = "Info"
"initiateSetup" = "Initiate Setup"
"inprogress" = "In Progress"
"insufficentFund" = "Insufficient funds"
"invalidAddress" = "Invalid address"
"invalidAmount" = "Invalid amount"
"invalidHex"     = "Invalid hex"
"invalidPassphrase" = "Password entered was not valid."
"invalidSeedPhrase" = "Invalid seed phrase"
"invalidSignature" = "Invalid signature or message"
"ipAddress" = "IP address"
"justNow" = "Just now"
"keepAppOpen" = "Keep this app opened"
"keepInMind" = "Keep in mind"
"key" = "Key"
"labelSpendable" = "Spendable"
"language" = "Language"
"lastBlockHeight" = "Last Block Height"
"latestBlock" = "Latest block"
"license" = "License"
"lifeSpan" = "Life Span"
"live" = "Live"
"liveIn" = "Live in"
"liveInfo" = "Waiting to be chosen to vote"
"liveInfoDisc" = "live for %v days, (%v of %v remaining)"
"liveInfoDiscSub" = "There is a 0.5% chance of expiring before being chosen to vote (this expiration returns the original Stake price without a reward)"
"liveTickets" = "Live Tickets"
"loading" = "Loading..."
"loadingPrice" = "Loading price"
"locked" = "Locked"
"lockedByTickets" = "Locked By Tickets"
"lockedin" = "Locked In"
"logLevel" = "Log Level"
"logLevelCritical"  = "Critical"
"logLevelDebug"  = "Debug"
"logLevelError"  = "Error"
"logLevelInfo"   = "Info"
"logLevelOff"    = "Off"
"logLevelTrace"  = "Trace"
"logLevelWarn"   = "Warn"
"manual" = "Manual"
"manualSetUp" = "Manual Setup"
"maturity" = "Maturity"
"max" = "MAX"
"message" = "Message"
"minimumAssetType" = "You need to create at least one DCR and one BTC wallet to use the exchange."
"minMax" = "Min: %f . Max: %f"
"mins" = "Mins"
"minuteAgo" = "%d minute ago"
"minutesAgo" = "%d minutes ago"
"missedOn" = "Missed on"
"missedTickets"="Missed Ticket"
"mix" = "Mix"
"mixed" = "Mixed"
"mixedAccDisabled" = "Receiving to mixed account is disabled by StakeShuffle settings to protect your privacy"
"mixedAccHidden" = "The Mixed Account is Hidden"
"mixedAccount" = "Mixed account"
"mixer" = "Mixer"
"mixerAccErrorMsg" = "There are existing accounts named mixed or unmixed. Please change the name to something else for now. You can change them back after the setup."
"mixerRunning" = "Mixer is running..."
"mixerShutdown" = "The mixer will automatically stop when unmixed balance are fully mixed."
"mixerStart" = "Mixer start Successfully"
"mixingActivity" = "Mixing Activities"
"monthAgo" = "%d month ago"
"monthsAgo" = "%d months ago"
"more" = "More"
"moveFundsFrmDefaultToUnmixed" = "Automatically move funds from default to unmixed account"
"moveToUnmixed" = "Move funds to unmixed account"
"multipleMixerAccNeeded" = "Set up mixer by creating two needed accounts"
"myAcct" = "My account"
"nConfirmations" = "%d Confirmations"
"network" = "Network"
"neverSynced" = "Never Synced"
"newest" = "Newest"
"newProposalUpdate" = "New update for proposal with Token: %s"
"newSpendingPassword" = "New spending passphrase"
"newStartupPass" = "New startup password"
"newWallet" = "New wallet"
"next" = "Next"
"no" = "No"
"noActiveTickets" = "No active tickets"
"noAgendaYet" = "No agendas yet"
"noConnectedPeer" = "no connected peers."
"noInternet" = "no Internet Connectivity."
"nonAccSelector" = "This widget isn't set to show accounts"
"none" = "None"
"noOrders" = "Orders you create will be shown here."
"noPoliciesYet" = "No policies yet"
"noProposal" = "No proposals %v"
"noReward" = "Stakey sees no rewards"
"notAllowed" = "%s API not allowed by current network settings."
"notApplicable" = "N/A"
"notAvailable" = "Not available"
"notBackedUp" = "Backup needed"
"notConnected" = "Not connected to decred network"
"note" = "Note"
"notEnoughVotes" = "You don't have enough votes"
"noTickets" = "No tickets yet"
"notifications" = "Notifications"
"notOwned" = "Valid address not owned by you."
"noTransactions" = "No transactions"
"notSameAccoutMixUnmix" = "Cannot use same account for mixed & unmixed"
"noUTXOs" = "No UTXOs Available"
"noValidAccountFound" = "no valid account found"
"noValidWalletFound" = "no valid wallet found"
"noVSPLoaded" = "No vsp loaded. Check internet connection and try again."
"noWalletLoaded" = "No wallet loaded"
"numberOfVotes" = "You have %d votes"
"offChainVote" = "Off-chain voting for development and marketing initiatives funded by the Decred treasury."
"offline" = "Offline, "
"ok" = "OK"
"oldest" = "Oldest"
"onChainVote" = "On-chain voting for upgrading the Decred network consensus rules."
"online" = "Online, "
"openingWallet" = "Opening wallets..."
"orderCreated" = "Order created successfully"
"orderDetails" = "Order Details"
"orderReceivingTo" = "To: %s (%s)"
"orderSendingFrom" = "From: %s (%s)"
"orderSettingsSaved" = "Order Settings saved!"
"orderSubmitted" = "Order Submitted"
"overview" = "Overview"
"owned" = "Valid address owned by you."
"pageWarningNotSync" = "Page cannot be accessed because the wallet is not synced, please sync your wallet and try again"
"pageWarningSync" = "Page cannot be accessed because the wallet sync is in progress, please wait for the sync to complete"
"passwordNotMatch" = "Passwords do not match"
"pasteSeedWords" = "Paste Seed Words"
"peer" = "Peer"
"peers" = "peers"
"peersConnected" = "Peers connected"
"pending" = "Pending"
"percentageMixed" = "%v%% Mixed"
"piKey" = "Pi key"
"policySetSuccessfully" = "Your treasury policy has been successfully updated!"
"priority" = "Priority%v"
"privacyInfo" = "%v When the mixer is activated, funds will be gradually transfered from the unmixed account to the mixed account. %v Important: keep this app open while mixer is running. %v The mixer routine will automatically stop when the unmixed balance is fully mixed.%v"
"privacyModeActive" = "(Network Privacy Is Enabled)"
"privacyModeInfo" = "Network Privacy Info"
"privacyModeInfoDesc" = "When enabled, all HTTP API calls are disabled, with the exception of Network Check API that is used to check if a wallet has internet access."
"privacySettings" = "Network Privacy"
"propFetching" = "Proposals fetching %s. %s"
"propNotif" = "Proposal notification"
"propNotification" = "Proposal notification %s"
"proposalAddedNotif" = "A new proposal has been added Name: %s"
"proposalInfo" = "Proposals and politeia notifications can be enabled or disabled from the settings page."
"proposals" = "Proposals"
"proposalVoteDetails" = "Proposal vote details"
"published" = "Published:   %s"
"published2" = "Published"
"purchased" = "Purchased"
"purchasedOn" = "Purchased On"
"purchasingAcct" = "Purchasing account"
"quorumRequirement" = "Quorum requirement:  %6.0f"
"rate" = "Rate"
"readyToMix" = "Ready to mix"
"rebroadcast" = "Rebroadcast"
"receive" = "Receive"
"received" = "Received"
"receiveInfo" = "To protect your privacy, a new address is generated each time you receive a payment."
"receiving" = "Receiving"
"receivingAddress" = "Receiving account"
"recentProposals" = "Recent Proposals"
"recentTransactions" = "Recent Transactions"
"reconnect" = "Reconnect"
"refresh" = "Refresh"
"rejected" = "Rejected"
"remove" = "Remove"
"removePeer" = "Remove specific peer"
"removePeerWarn" = "Are you sure you want to proceed with removing the specific peer?"
"removeUserAgent" = "Remove user agent"
"removeUserAgentWarn" = "Are you sure you want to proceed with removing the user agent?"
"removeWallet" = "Remove wallet from device"
"removeWalletInfo" = "%v Are you sure you want to remove %v %s%v? Enter the name of the wallet below to verify. %v"
"rename" = "Rename"
"renameAcct" = "Rename account"
"renameWalletSheetTitle" = "Rename wallet"
"republished" = "Republished unmined transactions to the %s network"
"rescan" = "Rescan"
"rescanBlockchain" = "Rescan blockchain"
"rescanInfo" = "Rescanning may help resolve some balance errors. This will take some time, as it scans the entire blockchain for transactions"
"rescanningBlocks" = "Rescanning blocks"
"rescanningHeaders" = "Rescanning headers · %v%%"
"rescanProgressNotification" = "Check progress in overview."
"restore" = "Restore"
"restoreExistingWallet" = "Restore existing wallet"
"restoreWallet" = "Restore wallet"
"restoreWithHex" = "Restore wallet using hex"
"resumeAccountDiscoveryTitle" = "Unlock to resume restoration"
"retry" = "Retry"
"revocation" = "Revocation"
"revoke" = "Revoke"
"revokeCause" = "Revocation cause"
"revoked" = "Revoked"
"revokeInfo" = "This Stake has been revoked."
"revokeInfoDisc" = "The Stake price will become spendable after %d blocks (~%s)"
"reward" = "Reward"
"rewardsEarned" = "Rewards Earned"
"save" = "Save"
"search" = "Search"
"secs" = "Secs"
"security" = "Security"
"securityTools" = "Security tools"
"securityToolsInfo" = "%v Various tools that help in different aspects of crypto currency security will be located here. %v"
"seeAll" = "See all"
"seedAlreadyExist" = "A wallet with an identical seed already exists."
"seedHex" = "Seed hex"
"seedPhraseToRestore" = "seed phrase is the only way to restore your wallet."
"seedPhraseVerified" = "Your seed phrase backup is verified"
"seedValidationFailed" = "Failed to verify. Please go through every wallet seed and try again."
"selectAcc" = "Select Account"
"selectAServer" = "Select A Server"
"selectAssetType" = "Select Asset Type"
"selectChangeAcc" = "%v Select the account you would like to use as your %v unmixed account. %v Note: you can compromise your privacy if you choose the wrong account %v"
"selectDexServerToOpen" = "Select the Dex server you would like to open."
"selectedAcct" = "Selected account: %s"
"selectedUTXO" = "Selected UTXO"
"selectMixedAcc" = "%v Select the account you would like to use as your %v mixed account. %v Note: you can compromise your privacy if you choose the wrong account %v"
"selectOption" = "Select one of the options below to vote"
"selectPhrasesToVerify" = "Select the correct phrases to verify."
"selectServerTitle" = "Select the exchange server you would like to use."
"selectTicket" = "Select a ticket to vote"
"selectUTXO" = "Select UTXO"
"selectVSP" = "Select a VSP..."
"selectWallet" = "Select Wallet"
"selectWalletToOpen" = "Select the wallet you would like to open."
"selectWalletType" = "Select the type of wallet you want to create"
"send" = "Send"
"sendConfModalTitle" = "You're about to send"
"sendInfo" = "Input or scan the destination wallet address and input the amount to send funds."
"sending" = "Sending"
"sendingAcct" = "Sending account"
"sendingFrom" = "Sending from"
"sendWarning" = "Your DCR will be sent after this step."
"sent" = "Sent"
"server" = "Server"
"setchoice" = "Set Choice"
"setGapLimit" = "Set Gap Limit"
"setGapLimitInfo" = "%v In some rare circumstances, address may not be discovered with the default gap limit of 20. It's recommended to only use this functionality after trying other options. And be aware that raising the gap limit above 100 will lead to excessive loading times to complete this request. %v"
"settings" = "Settings"
"setTreasuryPolicy" = "Set treasury policy"
"setUp" = "Set up"
"setupMixerInfo" = "%v Two dedicated accounts %v mixed %v & %v unmixed %v will be created in order to use the mixer. %v This action cannot be undone.%v"
"setUpNeededAccs" = "Set up needed accounts"
"setUpPrivacy" = "Using StakeShuffle increases the privacy of your wallet transactions."
"setUpStakeShuffle" = "Set up StakeShuffle"
"setupStartupPassword" = "Set up startup password"
"signature" = "Signature"
"signCopied" = "Signature copied"
"signMessage" = "Sign message"
"signMessageInfo" = "%v Signing a message with an address' private key allows you to prove that you are the owner of a given address to a possible counterparty.%v"
"source" = "Source"
"sourceModalInfo" = "Wallets that have not completed sync will be hidden from the list. %v Refunds and leftover change will be returned to the selected source account"
"sourceWalletNotSynced" = "Source wallet is not synced"
"spanish" = "Spanish"
"spendableIn" = "Spendable in"
"spendingPassword" = "Spending passphrase"
"spendingPasswordInfo" = "A spending password helps secure your wallet transactions."
"spendingPasswordInfo2" = "This spending password is for the new wallet only"
"spendingPasswordUpdated" = "Spending passphrase updated"
"stake" = "Stake"
"stakeAge" = "Stake age"
"staked" = "Staked"
"stakeShuffle" = "StakeShuffle"
"staking" = "Staking"
"stakingActivity" = "Staking Activities"
"startupPassConfirm" = "Startup password changed"
"startupPassword" = "Startup Password"
"startupPasswordEnabled" = "Startup password %v"
"startupPasswordInfo" = "Startup password helps protect your wallet from unauthorized access."
"statistics" = "Statistics"
"status" = "Status"
"step1" = "Step 1/2"
"step2of2" = "Step 2/2"
"submit" = "Submit"
"summary" = "Summary"
"sureToCancelMixer" = "Are you sure you want to cancel mixer action?"
"sureToExitBackup" = "Are you sure you want to exit the seed backup process?"
"sureToSafeStoreSeed" = "Be sure to store your seed phrase backup in a secure location."
"sync" = "Sync"
"syncCompTime" = "Est. Sync completion time"
"synced" = "Synced"
"syncingProgress" = "Syncing progress"
"syncingProgressStat" = "%s behind"
"syncingState" = "Syncing..."
"syncSteps" = "Step %d/3"
"takenAccount" = "Account name is taken"
"tapToCopy" = "(Tap to copy)"
"ticektVoted" =  "A ticket just voted\nVote reward: %s DCR"
"ticket" = "Ticket"
"ticketConfirmed" = "Ticket(s) Confirmed"
"ticketDetails" = "Ticket details"
"ticketError" = "Ticket buyer account error: %v"
"ticketPrice" = "Ticket Price"
"ticketRecord" = "Ticket Record"
"ticketRevoked" = "A ticket was revoked"
"ticketRevokedTitle" = "Ticket, Revoked"
"tickets" = "Tickets"
"ticketSettingSaved" = "Auto ticket purchase setting saved successfully."
"ticketVotedTitle" = "Ticket, Voted"
"timeLeft" = "%v left"
"to" = "To"
"token" = "Token:   %s"
"total" = "Total"
"totalAmount" = "Total Amount"
"totalBalance" = "Total Balance"
"totalCost" = "Total cost"
"totalVotes" = "Total votes:  %6.0f"
"totalVotesReverse" = "%d Total votes"
"transactionDetails" = "Transaction details"
"transactionId" = "Transaction ID"
"transactions" = "Transactions"
"transferred" = "Transferred"
"treasury" = "Treasury"
"treasurySpending" = "Treasury Spending"
"treasurySpendingInfo" = "Spending treasury funds now requires stakeholders to vote on the expenditure. You can participate and set a voting policy for treasury spending by a particular Governance Key. The keys can be verified in the dcrd source."
"txConfModalInfoTxt" = "<b>Unmixed accounts are hidden</b>. Spending from unmixed accounts is disabled by stakeshuffle settings to protect your privacy"
"txDetailsInfo" = "%v Tap on %v blue text %v to copy the item %v"
"txEstimateErr" = "Error estimating transaction: %v"
"txFee" = "Transaction Fee"
"txHashCopied" = "Transaction Hash copied"
"txNotification" = "Transaction Notification"
"txOverview" = "Transaction Overview"
"txSent" = "Transaction sent!"
"txSize" = "Transaction Size%v"
"txStatusPending"         = "Pending (%v of %v confirmations)" 
"type" = "Type"
"unconfirmedFunds" = "Allow spending unconfirmed funds"
"unconfirmedTx"    = "Unconfirmed"
"underReview" = "Under Review"
"unknown" = "Unknown"
"unlock" = "Unlock"
"unlockWithPassword" = "Unlock with password"
"unmined" = "Unmined"
"unminedInfo" = "Broadcasted %v"
"unmixed" = "Unmixed"
"unmixedAccount" = "Unmixed account"
"unmixedBalance" = "Unmixed balance"
"upcomming" = "Upcoming"
"updated" = "Updated"
"updatePreference" = "Update Preference"
"updateVotePref" = "Update Voting Preference"
"uptime" = "Uptime"
"usdBinance" = "USD (Binance)"
"usdBittrex" = "USD (Bittrex)"
"useMixer" = "How to use the mixer?"
"userAgent" = "User agent"
"userAgentDialogTitle" = "Set up user agent"
"userAgentSummary" = "For exchange rate fetching"
"validAddress" = "Valid address"
"validate" = "Validate"
"validateAddr" = "Validate address"
"validateHostErr" = "%s is not a valid IP or URL address"
"validateMsg" = "Validate address"
"validateNote" = "Enter an address to validate:"
"validateWalSeed" = "Validate wallet seeds"
"validSignature" = "Valid signature"
"verify" = "Verify"
"verifyGovernanceKeys" = "Verify Governance Keys"
"verifyMessage" = "Verify message"
"verifyMessageInfo" = "%v You can use this form to verify the signature's validity after you or your counterparty have generated one.%v After you've input the address, message, and signature, you'll see VALID if the signature matches the address and message correctly, and INVALID otherwise.%v"
"verifyMsgError" = "Error verifying message: %v"
"verifyMsgNote" = "Enter the address, signature, and message to verify:"
"verifySeed" = "Verify Seed Phrase"
"verifySeedInfo" = "Verify your seed phrase backup so you can recover your funds when needed."
"version" = "Version"
"viewAppLog" = "View Application Log"
"viewDetails" = "View details"
"viewOnExplorer" = "View on block explorer"
"viewOnPoliteia" = "View on Politeia"
"viewProperty" = "View property"
"viewSeedPhrase" = "View seed phrase"
"viewTicket" = "View associated ticket"
"vote" = "Vote"
"votechoice" = "Vote Choice"
"voteConfirm" = "Confirm to vote"
"voted" = "Voted"
"votedInfo" = "Congratulations! This Stake has voted."
"votedInfoDisc" = "The Stake price + reward will become spendable after %d blocks (~%s)"
"votedOn" = "Voted on"
"voteEndedNotif" = "Voting has ended for proposal with Token: %s"
"voteSent" = "Vote sent successfully, refreshing proposals!"
"voteStartedNotif" = "Voting has started for proposal with Token: %s"
"voteTooltip" = "%d %% Yes votes required for approval"
"voteUpdated" = "Vote choice updated successfully"
"voting" = "Voting"
"votingAuthority" = "Voting Authority"
"votingDashboard" = "Voting Dashboard"
"votingInProgress" = "Voting in progress..."
"votingPreference" = "Voting Preference:"
"votingServiceProvider" = "Voting service provider"
"votingWallet" = "Voting wallet"
"vsp" = "VSP"
"vspAPI" = "VSP API"
"vspFee" = "VSP Fee"
"waitingForAdmin" = "Waiting for admin to trigger the start of voting"
"waitingForAuthor" = "Waiting for author to authorize voting"
"waitingState" = "Waiting..."
"walletCreated" = "Wallet created"
"walletDirectory" = "Wallet data directory"
"walletExist" = "Wallet with name: %s already exist"
"walletLengthError" = "Wallet name must be less than 20 characters"
"walletLog" = "Wallet log"
"walletName" = "Wallet name"
"walletNameMismatch" = "Wallet name entered doesn't match selected one"
"walletNotExist" = "Wallet with ID: %v does not exist"
"walletNotSynced" = "Not Synced"
"walletRemoved" = "Wallet removed"
"walletRemoveInfo" = "Make sure to have the seed phrase backed up before removing the wallet"
"walletRenamed" = "Wallet renamed successfully"
"walletRestored" = "Wallet restored"
"walletRestoreMsg" = "You can restore this wallet from seed phrase after it is deleted."
"wallets" = "Wallets"
"walletsEnabledPrivacy" = "For wallets that have enabled privacy before."
"walletSettings" = "Wallet settings"
"walletStatus" = "Wallet Status:"
"walletSyncing" = "Wallet is syncing, please wait"
"walletToPurchaseFrom" = "Wallet to purchase from: %s"
"warningVote" = "You cannot vote with a watch only wallet"
"warningWatchWallet" = "You would be receiving to a read only wallet"
"watchOnly" = "Watch-Only"
"watchOnlyWalletImported" = "Watch only wallet imported"
"watchOnlyWalletRemoveInfo" = "The watch-only wallet will be removed from your app"
"watchOnlyWallets" = "Watch-only wallets"
"webURL" = "Web URL"
"weekAgo" = "%d week ago"
"weeksAgo" = "%d weeks ago"
"welcomeNote" = "Welcome to Cryptopower Wallet."
"whatToCallWallet" = "What would you like to call your wallet?"
"word" = "Word"
"writeDownAll33Words" = "Write down all 33 words in the correct order."
"writeDownSeed" = "Write down seed phrase"
"wroteAllWords" = "I have written down all 33 words"
"xInputsConsumed" = "%d Inputs consumed"
"xOutputCreated" = "%d Outputs created"
"xpubKeyErr" = "Error checking xpub: %v"
"xpubWalletExist" = "A wallet with an identical extended public key already exists."
"yearAgo" = "%d year ago"
"yearsAgo" = "%d years ago"  
"yes" = "Yes"
"yesterday" = "Yesterday"
"yourAddress" = "Your Address"
"yourSeedWord" = "Your 33-word seed phrase"
"yourself" = "Yourself"
"orderScheduler" = "Order Scheduler"
"start" = "Start"
`
