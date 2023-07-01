defmodule IsbnVerifier do
  @doc """
    Checks if a string is a valid ISBN-10 identifier

    ## Examples

      iex> IsbnVerifier.isbn?("3-598-21507-X")
      true

      iex> IsbnVerifier.isbn?("3-598-2K507-0")
      false

  """
  @spec isbn?(String.t()) :: boolean
  def isbn?(isbn) do
    isbn = String.replace(isbn, "-", "")
    if Regex.match?(~r/^\d{9}(\d|X)$/, isbn), do: valid_sum?(isbn), else: false
  end

  defp valid_sum?(isbn), do: rem(sum(isbn), 11) == 0

  defp sum(isbn) do
  isbn
  |> String.split("", trim: true)
  |> Enum.map(fn (x) -> if x == "X", do: 10, else: String.to_integer(x) end)
  |> Enum.zip(10..1)
  |> Enum.reduce(0, fn {a, b}, acc -> acc + a * b end)
  end
end
