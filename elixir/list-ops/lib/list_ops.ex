defmodule ListOps do
  # Please don't use any external modules (especially List or Enum) in your
  # implementation. The point of this exercise is to create these basic
  # functions yourself. You may use basic Kernel functions (like `Kernel.+/2`
  # for adding numbers), but please do not use Kernel functions for Lists like
  # `++`, `--`, `hd`, `tl`, `in`, and `length`.

  @spec count(list) :: non_neg_integer
  def count([]), do: 0
  def count([_head | tail]), do: 1 + count(tail)

  @spec reverse(list) :: list
  def reverse(list), do: reverse_list(list, [])
  defp reverse_list([], rev), do: rev
  defp reverse_list([head | tail], rev), do: reverse_list(tail, [head | rev])

  @spec map(list, (any -> any)) :: list
  def map([], _), do: []
  def map([head | tail], f), do: [f.(head) | map(tail, f)]

  @spec filter(list, (any -> as_boolean(term))) :: list
  def filter([], _), do: []
  def filter([head | tail], f) do
    cond do
      f.(head) -> [head | filter(tail, f)]
      true -> filter(tail, f)
    end
  end

  @type acc :: any
  @spec foldl(list, acc, (any, acc -> acc)) :: acc
  def foldl([], acc, _), do: acc
  def foldl([head | tail], acc, f), do: foldl(tail, f.(head, acc), f)

  @spec foldr(list, acc, (any, acc -> acc)) :: acc
  def foldr([], acc, _), do: acc
  def foldr([head | tail], acc, f), do: f.(head, foldr(tail, acc, f))

  @spec append(list, list) :: list
  def append(a, b), do: do_append(reverse(a), b)
  defp do_append([], b), do: b
  defp do_append([head | tail], b), do: do_append(tail, [head | b])

  @spec concat([[any]]) :: [any]
  def concat([]), do: []
  def concat([head | tail]), do: append(head, concat(tail))
end
