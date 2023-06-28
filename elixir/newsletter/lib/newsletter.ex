defmodule Newsletter do
  def read_emails(path) do
    path
    |> File.read!
    |> String.split("\n", trim: true)
  end

  def open_log(path), do: File.open!(path, [:write])

  def log_sent_email(pid, email), do: IO.puts(pid, email)

  def close_log(pid), do: File.close(pid)

  def send_newsletter(emails_path, log_path, send_fun) do
    log = open_log(log_path)
    emails_path
    |> read_emails
    |> Enum.each(fn email -> with :ok <- send_fun.(email), do: log_sent_email(log, email) end)
    close_log(log)
  end
end
