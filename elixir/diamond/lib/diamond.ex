defmodule Diamond do
  @doc """
  Given a letter, it prints a diamond starting with 'A',
  with the supplied letter at the widest point.
  """
  @spec build_shape(char) :: String.t()
  def build_shape(letter), do: construct_shape(letter, ?A)

  defp construct_shape(letter, letter), do: construct_line(letter, letter)
  defp construct_shape(letter, current) do
    construct_line(letter, current) <> construct_shape(letter, current + 1) <> construct_line(letter, current)
  end

  defp construct_line(letter, ?A) do
    side_space(letter, ?A) ++ [?A | side_space(letter, ?A)] ++ [?\n]
    |> List.to_string
  end
  defp construct_line(letter, current) do
    side_space(letter, current) ++ [current | middle_space(current)] ++ [current | side_space(letter, current)] ++ [?\n]
    |> List.to_string
  end

  defp side_space(letter, current), do: List.duplicate(?\s, letter - current)

  defp middle_space(current), do: List.duplicate(?\s, 2 * (current - ?A) - 1)
end
