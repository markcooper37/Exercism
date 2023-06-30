defmodule Say do
  @digits %{0 => "", 1 => "one", 2 => "two", 3 => "three", 4 => "four", 5 => "five", 6 => "six", 7 => "seven", 8 => "eight", 9 => "nine"}
  @tens %{0 => "", 2 => "twenty", 3 => "thirty", 4 => "forty", 5 => "fifty", 6 => "sixty", 7 => "seventy", 8 => "eighty", 9 => "ninety"}
  @teens %{0 => "ten", 1 => "eleven", 2 => "twelve", 3 => "thirteen", 4 => "fourteen", 5 => "fifteen", 6 => "sixteen", 7 => "seventeen", 8 => "eighteen", 9 => "nineteen"}

  @doc """
  Translate a positive integer into English.
  """
  @spec in_english(integer) :: {atom, String.t()}
  def in_english(number) when number < 0 or number > 999999999999, do: {:error, "number is out of range"}
  def in_english(0), do: {:ok, "zero"}
  def in_english(number) do
    {:ok, number
    |> Integer.digits
    |> Enum.reverse
    |> Enum.chunk_every(3)
    |> Enum.map(&Enum.reverse/1)
    |> Enum.zip(["", " thousand", " million", " billion"])
    |> Enum.reverse
    |> Enum.map(&construct_words/1)
    |> Enum.filter(&(&1 != ""))
    |> Enum.join(" ")}
  end

  defp construct_words({num, scale}) do
    if (words = construct_phrase(num)) == "", do: "", else: words <> scale
  end

  defp construct_phrase([0, tens, units]), do: tens_and_units([tens, units])
  defp construct_phrase([hundreds, tens, units]), do: "#{@digits[hundreds]} hundred" <> if (word = tens_and_units([tens, units])) != "", do: " #{word}", else: ""
  defp construct_phrase(num), do: tens_and_units(num)

  defp tens_and_units([1, units]), do: @teens[units]
  defp tens_and_units([0, units]), do: @digits[units]
  defp tens_and_units([tens, 0]), do: @tens[tens]
  defp tens_and_units([tens, units]), do: "#{@tens[tens]}-" <> @digits[units]
  defp tens_and_units([units]), do: @digits[units]
end
