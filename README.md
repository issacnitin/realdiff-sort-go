# RealDiff Go sort-stability demo

RealDiff runs the same tests on both sides of this pull request and reports the runtime values that changed.

## How it works

1. Check out the base and pull-request revisions.
2. Build both through RealDiff's stable Go source rewriter.
3. Run the same Go tests on both, recording observed function arguments and return values.
4. Diff those execution traces instead of inferring behavior from the source diff.

This is not mutation testing, static analysis, or coverage. No production code or test is mutated in the repository, RealDiff does not generate tests, and it observes only code this test suite executes.

## Worked example

The pull request chooses the lower-overhead unstable sort. In this block, `-` is the base, `+` is the proposal, and the single important change is `sort.SliceStable` becoming `sort.Slice`:

```diff
-sort.SliceStable(ordered, func(left, right int) bool {
+sort.Slice(ordered, func(left, right int) bool {
	 return ordered[left].Priority < ordered[right].Priority
 })
```

Both calls sort by priority, so the edit looks like a local performance refactor. The stable base preserves declaration order and selects `Z_CLEARANCE`. The unstable sort rearranges equal-priority entries and selects `A_SEASONAL` in this deterministic fixture.

The following block labels the exact values RealDiff observed before and after the edit:

```text
BASE  selectDiscount(100) -> Z_CLEARANCE
PR    selectDiscount(100) -> A_SEASONAL
BASE  checkoutTotal(100) -> (60, Z_CLEARANCE)
PR    checkoutTotal(100) -> (85, A_SEASONAL)
```

Neither pricing function is in the diff; only `src/config.go` changed. All three tests execute the path. The two broad total assertions still pass because 85 is discounted and does not exceed 100. The exact clearance-winner assertion is the only test written to check the changed selection.

## Why the finding is focused

RealDiff runs the base more than once and subtracts observations that disagree with themselves, removing timestamps, GUIDs, hash-order variation, and similar self-noise.

The changed rule affects its callers, but RealDiff collapses that propagation and reports the first changed behavior in unedited `src/pricing.go`.

## Run it

The command below runs the demo's three tests:

```bash
go test ./...
```
