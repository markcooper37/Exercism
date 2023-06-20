defmodule PerfectNumbers do
  @doc """
  Determine the aliquot sum of the given `number`, by summing all the factors
  of `number`, aside from `number` itself.

  Based on this sum, classify the number as:

  :perfect if the aliquot sum is equal to `number`
  :abundant if the aliquot sum is greater than `number`
  :deficient if the aliquot sum is less than `number`
  """
  @spec classify(number :: integer) :: {:ok, atom} | {:error, String.t()}
  def classify(number) when number < 1, do: {:error, "Classification is only possible for natural numbers."}
  def classify(number) do
    sum = factor_sum(number)
    cond do
      number == sum -> {:ok, :perfect}
      number < sum -> {:ok, :abundant}
      true -> {:ok, :deficient}
    end
  end

  defp factor_sum(number) do
    Enum.reduce(1..div(number, 2)//1, 0, fn (value, acc) -> if rem(number, value) == 0, do: acc + value, else: acc end)
  end
end
