defmodule CollatzConjecture do
  @doc """
  calc/1 takes an integer and returns the number of steps required to get the
  number to 1 when following the rules:
    - if number is odd, multiply with 3 and add 1
    - if number is even, divide by 2
  """
  @spec calc(input :: pos_integer()) :: non_neg_integer()
  def calc(input) when is_integer(input) and input > 0 do
    cond do
      input == 1 -> 0
      rem(input, 2) == 0 -> 1 + calc(trunc(input / 2))
      true -> 1 + calc(3 * input + 1)
    end
  end
end
