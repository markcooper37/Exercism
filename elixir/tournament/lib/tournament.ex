defmodule Tournament do
  @result_options ["win", "draw", "loss"]
  @header "Team                           | MP |  W |  D |  L |  P"

  @doc """
  Given `input` lines representing two teams and whether the first of them won,
  lost, or reached a draw, separated by semicolons, calculate the statistics
  for each team's number of games played, won, drawn, lost, and total points
  for the season, and return a nicely-formatted string table.

  A win earns a team 3 points, a draw earns 1 point, and a loss earns nothing.

  Order the outcome by most total points for the season, and settle ties by
  listing the teams in alphabetical order.
  """
  @spec tally(input :: list(String.t())) :: String.t()
  def tally([]), do: @header
  def tally(input) do
    compile_results(%{}, input)
    |> Map.to_list
    |> Enum.sort(&compare/2)
    |> Enum.map(&stringify_results/1)
    |> Enum.join("\n")
    |> add_header
  end

  defp compile_results(teams, []), do: teams
  defp compile_results(teams, [result | results]) do
    case String.split(result, ";") do
      [team_1, team_2, result] when result in @result_options -> compile_results(update_results(teams, [team_1, team_2, result]), results)
      _ -> compile_results(teams, results)
    end
  end

  defp update_results(results, [team_1, team_2, result]) do
    results = Map.put_new(results, team_1, %{wins: 0, draws: 0, losses: 0})
    results = Map.put_new(results, team_2, %{wins: 0, draws: 0, losses: 0})
    case result do
      "win" -> (
        results
        |> add_result(team_1, :wins)
        |> add_result(team_2, :losses)
      )
      "draw" -> (
        results
        |> add_result(team_1, :draws)
        |> add_result(team_2, :draws)
      )
      "loss" -> (
        results
        |> add_result(team_1, :losses)
        |> add_result(team_2, :wins)
      )
    end
  end

  defp add_result(results, team, result) do
    Map.put(results, team, Map.update!(results[team], result, &(&1 + 1)))
  end

  defp compare({name_1, results_1}, {name_2, results_2}) do
    if points(results_1) == points(results_2), do: name_1 <= name_2, else: points(results_1) >= points(results_2)
  end

  defp stringify_results({name, results}) do
    "#{String.pad_trailing(name, 30)} | #{format_result_value(matches_played(results))} | #{format_result_value(results[:wins])} | #{format_result_value(results[:draws])} | #{format_result_value(results[:losses])} | #{format_result_value(points(results))}"
  end

  defp format_result_value(value), do: String.pad_leading(Integer.to_string(value), 2)

  defp matches_played(results), do: results[:wins] + results[:draws] + results[:losses]

  defp points(results), do: 3 * results[:wins] + results[:draws]

  defp add_header(results_string) do
    @header <> "\n" <> results_string
  end
end
