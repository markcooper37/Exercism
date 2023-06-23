defmodule Yacht do
  @type category ::
          :ones
          | :twos
          | :threes
          | :fours
          | :fives
          | :sixes
          | :full_house
          | :four_of_a_kind
          | :little_straight
          | :big_straight
          | :choice
          | :yacht

  @doc """
  Calculate the score of 5 dice using the given category's scoring method.
  """
  @spec score(category :: category(), dice :: [integer]) :: integer
  def score(:ones, dice), do: Enum.count(dice, &(&1 == 1))
  def score(:twos, dice), do: 2 * Enum.count(dice, &(&1 == 2))
  def score(:threes, dice), do: 3 * Enum.count(dice, &(&1 == 3))
  def score(:fours, dice), do: 4 * Enum.count(dice, &(&1 == 4))
  def score(:fives, dice), do: 5 * Enum.count(dice, &(&1 == 5))
  def score(:sixes, dice), do: 6 * Enum.count(dice, &(&1 == 6))
  def score(:choice, dice), do: Enum.sum(dice)
  def score(:yacht, [a, a, a, a, a]), do: 50
  def score(:yacht, _), do: 0
  def score(category, dice) do
    case {category, Enum.sort(dice)} do
      {:little_straight, [1, 2, 3, 4, 5]} -> 30
      {:big_straight, [2, 3, 4, 5, 6]} -> 30
      {:full_house, [a, a, a, a, a]} -> 0
      {:full_house, [a, a, a, b, b]} -> 3 * a + 2 * b
      {:full_house, [a, a, b, b, b]} -> 2 * a + 3 * b
      {:four_of_a_kind, [a, a, a, a, _]} -> 4 * a
      {:four_of_a_kind, [_, a, a, a, a]} -> 4 * a
      _ -> 0
    end
  end
end
