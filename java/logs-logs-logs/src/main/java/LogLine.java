public class LogLine {

    private LogLevel logLevel;
    private String message;

    public LogLine(String logLine) {
        switch (logLine.substring(1, 4)) {
            case "TRC":
                this.logLevel = LogLevel.TRACE;
                break;
            case "DBG":
                this.logLevel = LogLevel.DEBUG;
                break;
            case "INF":
                this.logLevel = LogLevel.INFO;
                break;
            case "WRN":
                this.logLevel = LogLevel.WARNING;
                break;
            case "ERR":
                this.logLevel = LogLevel.ERROR;
                break;
            case "FTL":
                this.logLevel = LogLevel.FATAL;
                break;
            default:
                this.logLevel = LogLevel.UNKNOWN;
        }
        this.message = logLine.substring(7);
    }

    public LogLevel getLogLevel() {
        return this.logLevel;
    }

    public String getOutputForShortLog() {
        return this.logLevel.getNumber() + ":" + this.message;
    }
}
