defmodule MatchingBrackets do

  @doc """
  Checks that all the brackets and braces in the string are matched correctly, and nested correctly
  """
  @spec check_brackets(String.t()) :: boolean
  def check_brackets(str) do
    String.codepoints(str)
    |> do_check
  end

  defp do_check(characters, brackets \\ [])
  defp do_check([], brackets), do: length(brackets) == 0
  defp do_check([char | chars], brackets) when char in ["[", "{", "("] do
    do_check(chars, [char | brackets])
  end
  defp do_check([char | _], []) when char in ["]", "}", ")"], do: false
  defp do_check([char | chars], [bracket | brackets]) when char in ["]", "}", ")"] do
    if bracket_matches?(bracket, char), do: do_check(chars, brackets), else: false 
  end
  defp do_check([_ | chars], brackets), do: do_check(chars, brackets)

  defp bracket_matches?(left, "]"), do: left == "["
  defp bracket_matches?(left, "}"), do: left == "{"
  defp bracket_matches?(left, ")"), do: left == "("
end
