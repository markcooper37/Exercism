defmodule TopSecret do
  def to_ast(string) do
    Code.string_to_quoted!(string)
  end

  def decode_secret_message_part(ast, acc) do
    case ast do
      {op, _, [{:when, _, [{name, _, args} | _]} | _]} when op == :def or op == :defp -> {ast, [String.slice(Atom.to_string(name), 0, find_length(args)) | acc]}
      {op, _, [{name, _, args} | _]} when op == :def or op == :defp -> {ast, [String.slice(Atom.to_string(name), 0, find_length(args)) | acc]}
      _ -> {ast, acc}
    end
  end

  defp find_length(args) do
    if args == nil, do: 0, else: Enum.count(args)
  end

  def decode_secret_message(string) do
    {_, message} = Macro.prewalk(to_ast(string), [], &decode_secret_message_part/2)
    message
    |> Enum.reverse
    |> Enum.join
  end
end
