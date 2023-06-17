defmodule TakeANumber do
  def start() do
    spawn(fn -> loop(0) end)
  end

  def loop(value) do
    receive do 
      {:report_state, sender_pid} -> (
        send(sender_pid, value)
        loop(value)
      )
      {:take_a_number, sender_pid} -> (
        send(sender_pid, value + 1)
        loop(value + 1)
      )
      :stop -> nil
      _ -> loop(value)
    end
  end
end
