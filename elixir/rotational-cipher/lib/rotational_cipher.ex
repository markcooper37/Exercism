defmodule RotationalCipher do
  @doc """
  Given a plaintext and amount to shift by, return a rotated string.

  Example:
  iex> RotationalCipher.rotate("Attack at dawn", 13)
  "Nggnpx ng qnja"
  """
  @spec rotate(text :: String.t(), shift :: integer) :: String.t()
  def rotate(text, shift) do
    text
    |> String.to_charlist
    |> Enum.map(&map_char(&1, shift))
    |> List.to_string
  end

  defp map_char(character, shift) do
    cond do
      character in ?A..?Z -> ?A + rem(character - ?A + shift, 26)
      character in ?a..?z -> ?a + rem(character - ?a + shift, 26)
      true -> character
    end
  end
end
