defmodule Strain do
  @doc """
  Given a `list` of items and a function `fun`, return the list of items where
  `fun` returns true.
  """
  @spec keep(list :: list(any), fun :: (any -> boolean)) :: list(any)
  def keep(list, fun) do
    do_keep(list, fun, [])
  end

  defp do_keep([], _fun, values), do: values
  defp do_keep([head | tail], fun, values) do
    cond do
      fun.(head) -> do_keep(tail, fun, values ++ [head])
      true -> do_keep(tail, fun, values)
    end
  end

  @doc """
  Given a `list` of items and a function `fun`, return the list of items where
  `fun` returns false.
  """
  @spec discard(list :: list(any), fun :: (any -> boolean)) :: list(any)
  def discard(list, fun) do
    do_discard(list, fun, [])
  end

  defp do_discard([], _fun, values), do: values
  defp do_discard([head | tail], fun, values) do
    cond do
      !(fun.(head)) -> do_discard(tail, fun, values ++ [head])
      true -> do_discard(tail, fun, values)
    end
  end
end
