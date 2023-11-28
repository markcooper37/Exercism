if (.series | length) == 0 then ("series cannot be empty" | halt_error)
elif .sliceLength < 0 then ("slice length cannot be negative" | halt_error)
elif .sliceLength == 0 then ("slice length cannot be zero" | halt_error)
elif (.series | length) < .sliceLength then ("slice length cannot be greater than series length" | halt_error)
else . as $input | [range(0;$input.series | length - $input.sliceLength + 1)] | map(. as $index | $input.series | .[$index:$index + $input.sliceLength]) end
