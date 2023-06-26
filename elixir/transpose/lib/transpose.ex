defmodule Transpose do
  @doc """
  Given an input text, output it transposed.

  Rows become columns and columns become rows. See https://en.wikipedia.org/wiki/Transpose.

  If the input has rows of different lengths, this is to be solved as follows:
    * Pad to the left with spaces.
    * Don't pad to the right.

  ## Examples

    iex> Transpose.transpose("ABC\\nDE")
    "AD\\nBE\\nC"

    iex> Transpose.transpose("AB\\nDEF")
    "AD\\nBE\\n F"
  """

  @spec transpose(String.t()) :: String.t()
  def transpose(input) do
    input
    |> String.split("\n")
    |> Enum.map(&String.split(&1, "", trim: true))
    |> create_transpose
    |> Enum.join("\n")
  end

  defp create_transpose(split) do
    create_lines(split, List.duplicate("", length(Enum.max_by(split, &length/1))), 0)
  end

  defp create_lines([], new_lines, _), do: new_lines
  defp create_lines([old_line | old_lines], new_lines, length) do
    create_lines(old_lines, add_line(new_lines, old_line, length), length + 1)
  end

  defp add_line(lines, [], _), do: lines
  defp add_line([new_line | new_lines], [old_char | old_chars], length) do
    [String.pad_trailing(new_line, length) <> old_char | add_line(new_lines, old_chars, length)]
  end
end
