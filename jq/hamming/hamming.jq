if (.strand1 | length) != (.strand2 | length) then ("strands must be of equal length" | halt_error)
else (.strand1 | split("")) as $strand1 | (.strand2 | split("")) as $strand2 | [range(0;$strand1 | length)] | map(select(. as $index | (($strand1 | .[$index])) != ($strand2 | .[$index]))) | length end
