# Contributing

## Contents
- [MR/Issues Naming Standard Format](#pr/issues-naming-standard-format)
- [When to Search the Issue Tracker](#when-to-search-the-issue-tracker)
- [When to Submit a New Issue](#when-to-submit-a-new-issue)
- [When to Submit a Merge Request](#when-to-submit-a-merge-request)
- [How to Contribute](#how-to-contribute)
  - [Suggested Toolkit](#suggested-toolkit)
  - [Git Workflow](#git-workflow)
- [Go Development Tips](#go-development-tips)

## Guidelines

### MR/Issues Naming Standard Format

For efficient records keeping and retrieval, follow the merge-request template.

### When to Search the Issue Tracker

If _any_ of the following apply:

- You believe you found a bug. It might already be logged, or even have a fix underway!
- You have an idea for an enhancement or new feature. Before you start coding away and submit a MR for your work, consult the issue tracker.
- You just want to find something to work on.

### When to Submit a New Issue

New issues may be submitted for enhancement requests as well as bug reports. However, we ask that you _please_ search first for similar existing issues to avoid posting a duplicate.

You are strongly advised to submit a new issue when you plan to perform work and submit a merge request (MR). See [When to Submit a Merge Request](#when-to-submit-a-merge-request) below.

A related matter of GitLab etiquette is when and how to post comments on Issues or MRs. Instead of simply posting "mee to! plus one", you can use the emoji responses to give a +1 or thumbs up.  Feel free to comment if you have more to add to the conversation.

### When to Submit a Merge Request

Before submitting a MR, check the issue tracker for existing issues or relevant discussion. See what has been done, if anything. Perhaps there is good reason why certain changes have not already been made.

If the planned commits will involve significant effort on your part, you definitely want to either (1) submit a new issue, or (2) announce your intention to work on an existing issue. Why? Someone else could already be working on the problem. Duplicating work really sucks. Also, there may be good reason why the change is not appropriate at the time. The best way to check is to head to the issue tracker.

Only submit a MR once the intended edits are either done or nearing completion. When submitting a MR with incomplete work, add the wip label.

**Do not submit a new MR if you have one In progress or Under Review**

## How to Contribute

### Suggested Toolkit

- Go language distribution - latest release or latest-1 (must be >=1.12). [download](https://golang.org/doc/install)
- git client with command line support. [download](https://git-scm.com/downloads)
- [GitLab](https://github.com/) account
- Visual Studio Code with Go extension plus `golangci-lint`
### Git Workflow

1. Fork the repository on GitLab.  Just click the little Fork button at https://github.com/crypto-power/cryptopower


2. Clone your newly forked cryptopower repository

```sh
git clone git@github.com:my-user-name/cryptopower.git
```

###### recommended

Setting your master branch to track this repository makes keeping everything up-to-date a breeze.
The rest of this workflow guide will assume that you have completed this step.

```sh
git remote add upstream https://github.com/crypto-power/cryptopower.git
git fetch upstream
git branch -u upstream/master master
```

3. Make a branch for your planned work, based on `master`

```sh
git checkout -b my-great-stuff master
```

4. Make edits. Review changes:

```sh
git status
git diff
```

5. Commit your work

```sh
# pick files you modified
git add -u
# don't forget to add that new file you made too
git add newfile.go
# one more check
git status
# make the commit
git commit # type a good commit message
```

6. Bring master up-to-date and rebase

Since the Cryptopower repo may have changes that you do not have locally, you'll want to merge in any changes and rebase.
Read [this](https://www.atlassian.com/git/tutorials/rewriting-history/git-rebase) if you need a primer on rebasing.

```sh
git checkout master
git merge
git checkout my-great-stuff
git rebase -i master
```

In the text editor, change the command from `pick` to `fixup` or `squash` for **all but the top** commit. Use `fixup` for little touchups to discard the commit's log message. If you want to use a different
commit message for everything, change the command from `pick` to `reword` on the top commit.
It should look something like this before saving.


![alt text](https://i.imgur.com/fOtaYtb.png "Rebase commmit command guide")


7. **If you have conflicts**, resolve them by iterating through the diffs one conflicting commit at a time.

```sh
# resolve conflicts
git add file1.go, file2.go, ...
git rebase --continue
# repeat until rebase completes
```

8. Push your branch to GitLab

Assuming `myremote` is the name of the remote used for *your* repository (by default, git created an alias, `origin`, for *your* forked repository in step 2 above, but you can [name it whatever you'd like](https://git-scm.com/book/en/v2/Git-Basics-Working-with-Remotes)):

```sh
git push -u myremote my-great-stuff
```

9. Create the merge request

On GitLab, select your branch in the dropdown menu (1) and click (2) to start a merge request.

![alt text](https://i.imgur.com/GXZTyiq.png "Merge Request submission guide")

Always:

- Type a detailed comment for the changes you are proposing.  Include motivation and a description of the code change.
- Highlight any breaking changes.  This includes any syntax changes, added or removed struct fields, interface changes, file renames or deletions, etc.
- Scroll down and review the code diffs. Verify that the changes are what you expect to see based on your earlier review of the diffs and your git commit log (you did that, right?).

Excellent [MR guidelines](https://github.com/kubernetes/community/blob/master/contributors/guide/merge-requests.md) from Kubernetes project.

10. Receive feedback and make changes

You will typically receive feedback from other developers. Before responding, take a moment to review the
[Code Contribution Guidelines](https://github.com/decred/dcrd/blob/master/docs/code_contribution_guidelines.md).

Work through the comments and resolve any confusion with others. Make whatever revisions are necessary.

11. Resubmitting

Commit your work as in step 5 above.

###### a)

Before resubmitting, clean up any little touchup commits you've made since the last time you pushed.
If you've only made one commit since then, you can skip this step.
For example, if you have made 3 commits since your last push, then run the following to "squash" them together.

```sh
git rebase -i HEAD~3
```

The number after the tilda (~) is the number of commits that you want to combine, including the one you did at the beginning of this step. Try not to squash post-review commits with pre-review commits. Leaving them separate makes navigating the changes easier.

###### b)

Then rebase the entire branch back to an updated master.

```sh
git checkout master
git merge
git checkout my-great-stuff
git rebase master
```

Note that the 4th command is different than step 6. You already performed the squash in the last step, so an interactive rebase (`-i`) is not needed here.

###### c)

Push the changes to your remote fork.

```sh
git push myremote my-great-stuff
```

Depending on what has changed, you will likely receive an error message rejecting your push for a misaligned branch tip. This is normal.
Rerun with the `--force` flag.

As soon as you push, your changes will be ready for review. There is typically no need to notify anybody that the changes have been made. GitLab takes care of that. Feel free to leave a comment on the merge request with a brief description of your changes.

## Go Development Tips

Use code linters. `golangci-lint` is suggested to run the whole lot.

Always `go fmt` your code before committing.

Read [Effective Go](https://golang.org/doc/effective_go.html).


### Further reading

[Code Contribution Guidelines](https://github.com/decred/dcrd/blob/master/docs/code_contribution_guidelines.md) (dcrd repo)