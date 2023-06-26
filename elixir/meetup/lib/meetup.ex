defmodule Meetup do
  @moduledoc """
  Calculate meetup dates.
  """

  @type weekday ::
          :monday
          | :tuesday
          | :wednesday
          | :thursday
          | :friday
          | :saturday
          | :sunday

  @type schedule :: :first | :second | :third | :fourth | :last | :teenth

  @weekday_map %{monday: 1, tuesday: 2, wednesday: 3, thursday: 4, friday: 5, saturday: 6, sunday: 7}

  @doc """
  Calculate a meetup date.

  The schedule is in which week (1..4, last or "teenth") the meetup date should
  fall.
  """
  @spec meetup(pos_integer, pos_integer, weekday, schedule) :: Date.t()
  def meetup(year, month, weekday, schedule) do
      Date.add(first_date(year, month, schedule), days_after_first_date(year, month, weekday, schedule))
  end

  defp days_after_first_date(year, month, weekday, schedule) do
    rem(@weekday_map[weekday] + 7 - Date.day_of_week(first_date(year, month, schedule)), 7)
  end

  defp first_date(year, month, :first), do: Date.new!(year, month, 1)
  defp first_date(year, month, :second), do: Date.new!(year, month, 8)
  defp first_date(year, month, :third), do: Date.new!(year, month, 15)
  defp first_date(year, month, :fourth), do: Date.new!(year, month, 22)
  defp first_date(year, month, :teenth), do: Date.new!(year, month, 13)
  defp first_date(year, month, :last) do
      Date.new!(year, month, Date.days_in_month(Date.new!(year, month, 1)) - 6)
  end
end
