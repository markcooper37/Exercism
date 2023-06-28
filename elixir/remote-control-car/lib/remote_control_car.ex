defmodule RemoteControlCar do
  @enforce_keys [:nickname]
  defstruct [:nickname, battery_percentage: 100, distance_driven_in_meters: 0]

  def new(nickname \\ "none"), do: %RemoteControlCar{nickname: nickname}

  def display_distance(remote_car = %RemoteControlCar{}), do: "#{remote_car.distance_driven_in_meters} meters"

  def display_battery(remote_car = %RemoteControlCar{}) do
    case battery = remote_car.battery_percentage do
      0 -> "Battery empty"
      _ -> "Battery at #{battery}%"
    end
  end

  def drive(remote_car = %RemoteControlCar{battery_percentage: 0}), do: remote_car
  def drive(remote_car = %RemoteControlCar{}) do
    %{remote_car | battery_percentage: remote_car.battery_percentage - 1, distance_driven_in_meters: remote_car.distance_driven_in_meters + 20}
  end
end
