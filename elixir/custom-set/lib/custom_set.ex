defmodule CustomSet do
  @opaque t :: %__MODULE__{map: map}

  defstruct map: %{}

  @spec new(Enum.t()) :: t
  def new(enumerable) do
    enumerable |> Enum.reduce(%__MODULE__{}, fn x, acc -> add(acc, x) end)
  end

  @spec empty?(t) :: boolean
  def empty?(custom_set) do
    Enum.count(custom_set.map) == 0
  end

  @spec contains?(t, any) :: boolean
  def contains?(custom_set, element) do
    Map.has_key?(custom_set.map, element)
  end

  @spec subset?(t, t) :: boolean
  def subset?(custom_set_1, custom_set_2) do
    Enum.all?(custom_set_1.map, fn {elem, _} -> contains?(custom_set_2, elem) end)
  end

  @spec disjoint?(t, t) :: boolean
  def disjoint?(custom_set_1, custom_set_2) do
    !Enum.any?(custom_set_1.map, fn {elem, _} -> contains?(custom_set_2, elem) end)
  end

  @spec equal?(t, t) :: boolean
  def equal?(custom_set_1, custom_set_2) do
    subset?(custom_set_1, custom_set_2) && subset?(custom_set_2, custom_set_1)
  end

  @spec add(t, any) :: t
  def add(%__MODULE__{map: m}, element) do
    %__MODULE__{map: Map.put(m, element, 1)}
  end

  @spec intersection(t, t) :: t
  def intersection(custom_set_1, custom_set_2) do
    custom_set_1.map
    |> Map.keys
    |> Enum.filter(&contains?(custom_set_2, &1))
    |> new
  end

  @spec difference(t, t) :: t
  def difference(custom_set_1, custom_set_2) do
    custom_set_1.map
    |> Map.keys
    |> Enum.reject(&contains?(custom_set_2, &1))
    |> new
  end

  @spec union(t, t) :: t
  def union(custom_set_1, custom_set_2) do
    custom_set_1.map
    |> Map.merge(custom_set_2.map)
    |> Map.keys
    |> new
  end
end
