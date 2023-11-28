def update_digit:
  . * 2 | if . > 9 then . - 9 else . end
;

. | explode | map(select(. != 32)) | if any(. < 48 or . > 57) then false
elif . | length | . == 1 then false 
else map(. - 48) | (. | length) as $length | [[range(0,$length)], .] | transpose | map(if (. | first | . + $length | . % 2 | . == 0) then (. | last | update_digit) else (. | last) end) | add | . % 10 | . == 0 end


