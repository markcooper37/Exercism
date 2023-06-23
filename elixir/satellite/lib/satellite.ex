defmodule Satellite do
  @typedoc """
  A tree, which can be empty, or made from a left branch, a node and a right branch
  """
  @type tree :: {} | {tree, any, tree}

  @doc """
  Build a tree from the elements given in a pre-order and in-order style
  """
  @spec build_tree(preorder :: [any], inorder :: [any]) :: {:ok, tree} | {:error, String.t()}
  def build_tree(preorder, inorder) do
    cond do
      length(preorder) != length(inorder) -> {:error, "traversals must have the same length"}
      preorder -- inorder != [] -> {:error, "traversals must have the same elements"}
      Enum.uniq(preorder) != preorder -> {:error, "traversals must contain unique items"}
      true -> {:ok, do_build_tree(preorder, inorder)}
    end
  end

  defp do_build_tree([], []), do: {}
  defp do_build_tree(preorder = [head | _tail], inorder) do
    {left_in, [_ | right_in]} = Enum.split_while(inorder, &(&1 != head))
    left_pre = Enum.filter(preorder, &(&1 in left_in))
    right_pre = Enum.filter(preorder, &(&1 in right_in))
    {do_build_tree(left_pre, left_in), head, do_build_tree(right_pre, right_in)}
  end
end
