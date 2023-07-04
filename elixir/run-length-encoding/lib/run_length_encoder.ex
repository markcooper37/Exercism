defmodule RunLengthEncoder do
  @doc """
  Generates a string where consecutive elements are represented as a data value and count.
  "AABBBCCCC" => "2A3B4C"
  For this example, assume all input are strings, that are all uppercase letters.
  It should also be able to reconstruct the data into its original form.
  "2A3B4C" => "AABBBCCCC"
  """
  @spec encode(String.t()) :: String.t()
  def encode(string) do
    Regex.split(~r/(.)\1{0,}/, string, include_captures: true, trim: true)
    |> Enum.map_join(&find_encoding/1)
  end

  defp find_encoding(string) do
    "#{if (len = String.length(string)) > 1, do: len}#{String.at(string, 0)}"
  end

  @spec decode(String.t()) :: String.t()
  def decode(string) do
    Regex.split(~r/\d*./, string, include_captures: true, trim: true)
    |> Enum.map_join(&find_decoding/1)
  end

  defp find_decoding(string) do
    {count, char} = String.split_at(string, String.length(string) - 1)
    len = if count == "", do: 1, else: String.to_integer(count)
    String.duplicate(char, len)
  end
end
