defmodule Atbash do
  @doc """
  Encode a given plaintext to the corresponding ciphertext

  ## Examples

  iex> Atbash.encode("completely insecure")
  "xlnko vgvob rmhvx fiv"
  """
  @spec encode(String.t()) :: String.t()
  def encode(plaintext) do
    plaintext
    |> String.downcase
    |> String.replace(~r/[\p{P}\s]/, "")
    |> String.to_charlist
    |> Enum.map(&map_char/1)
    |> Enum.chunk_every(5)
    |> Enum.join(" ")
  end

  @spec decode(String.t()) :: String.t()
  def decode(cipher) do
    cipher
    |> String.replace(~r/\s/, "")
    |> String.to_charlist
    |> Enum.map(&map_char/1)
    |> List.to_string
  end

  defp map_char(char) when char in ?a..?z, do: ?a + ?z - char
  defp map_char(char), do: char
end
