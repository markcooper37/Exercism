# The input will be null or an object that _may_ contain keys
#   actual_minutes_in_oven,
#   number_of_layers
#
# If the needed key is missing, use a default value:
#   zero minutes in oven,
#   one layer.
#
# Task: output a JSON object with keys:

40 as $expected | (.actual_minutes_in_oven // 0 | $expected - .) as $remaining | (.number_of_layers // 1 | . * 2) as $preparation | (.actual_minutes_in_oven + $preparation) as $total | 

{
  "expected_minutes_in_oven": $expected,
  "remaining_minutes_in_oven": $remaining,
  "preparation_time": $preparation,
  "total_time": $total
}
