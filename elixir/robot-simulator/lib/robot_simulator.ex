defmodule RobotSimulator do
  defstruct [:direction, :position]

  @turn_left [north: :west, west: :south, south: :east, east: :north]
  @turn_right [north: :east, east: :south, south: :west, west: :north]

  @type robot() :: any()
  @type direction() :: :north | :east | :south | :west
  @type position() :: {integer(), integer()}

  @doc """
  Create a Robot Simulator given an initial direction and position.

  Valid directions are: `:north`, `:east`, `:south`, `:west`
  """
  @spec create(direction, position) :: robot() | {:error, String.t()}
  def create(direction \\ :north, position \\ {0, 0}) do
    cond do
      !valid_direction?(direction) -> {:error, "invalid direction"}
      !valid_position?(position) -> {:error, "invalid position"}
      true -> %RobotSimulator{direction: direction, position: position}
    end
  end

  defp valid_position?({x1, x2}) do
    is_integer(x1) and is_integer(x2)
  end
  defp valid_position?(_), do: false

  defp valid_direction?(direction) do
    direction in [:north, :east, :south, :west]
  end

  @doc """
  Simulate the robot's movement given a string of instructions.

  Valid instructions are: "R" (turn right), "L", (turn left), and "A" (advance)
  """
  @spec simulate(robot, instructions :: String.t()) :: robot() | {:error, String.t()}
  def simulate(robot, ""), do: robot
  def simulate(robot, instructions) do
    {instruction, instructions} = String.split_at(instructions, 1)
    cond do
      !valid_instruction?(instruction) -> {:error, "invalid instruction"}
      true -> simulate(perform_action(robot, instruction), instructions)
    end
  end

  defp valid_instruction?(instruction) do
    instruction in ["L", "R", "A"]
  end

  defp perform_action(robot, "L"), do: %{robot | direction: @turn_left[robot.direction]}
  defp perform_action(robot, "R"), do: %{robot | direction: @turn_right[robot.direction]}
  defp perform_action(robot, "A"), do: advance(robot)

  defp advance(robot = %{direction: direction, position: {x1, x2}}) do
    case direction do
      :north -> %{robot | position: {x1, x2 + 1}}
      :east -> %{robot | position: {x1 + 1, x2}}
      :south -> %{robot | position: {x1, x2 - 1}}
      :west -> %{robot | position: {x1 - 1, x2}}
    end
  end

  @doc """
  Return the robot's direction.

  Valid directions are: `:north`, `:east`, `:south`, `:west`
  """
  @spec direction(robot) :: direction()
  def direction(robot), do: robot.direction

  @doc """
  Return the robot's position.
  """
  @spec position(robot) :: position()
  def position(robot), do: robot.position
end
