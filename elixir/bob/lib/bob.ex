defmodule Bob do
  @spec hey(String.t()) :: String.t()
  def hey(input) do
    cond do
      silence?(input) -> "Fine. Be that way!"
      yelling?(input) and question?(input) -> "Calm down, I know what I'm doing!"
      yelling?(input) -> "Whoa, chill out!"
      question?(input) -> "Sure."
      true -> "Whatever."
    end
  end

  defp silence?(input) do
    Regex.match?(~r/^\s*$/, input)
  end

  defp question?(input) do
    input
    |> String.trim_trailing
    |> String.ends_with?("?")
  end

  defp yelling?(input) do
    input == String.upcase(input) and Regex.match?(~r/\p{L}/, input)
  end
end
