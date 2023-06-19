defmodule LibraryFees do
  def datetime_from_string(string) do
    NaiveDateTime.from_iso8601!(string)
  end

  def before_noon?(datetime) do
    Time.compare(NaiveDateTime.to_time(datetime), ~T[12:00:00]) == :lt
  end

  def return_date(checkout_datetime) do
    days = if before_noon?(checkout_datetime), do: 28, else: 29
    NaiveDateTime.to_date(NaiveDateTime.add(checkout_datetime, days * 24 * 60 * 60))
  end

  def days_late(planned_return_date, actual_return_datetime) do
    difference = Date.diff(NaiveDateTime.to_date(actual_return_datetime), planned_return_date)
    if difference <= 0, do: 0, else: difference
  end

  def monday?(datetime) do
    Date.day_of_week(NaiveDateTime.to_date(datetime)) == 1
  end

  def calculate_late_fee(checkout, return, rate) do
    checkout = datetime_from_string(checkout)
    return = datetime_from_string(return)
    late = days_late(return_date(checkout), return)
    if monday?(return), do: div(rate * late, 2), else: rate * late
  end
end
