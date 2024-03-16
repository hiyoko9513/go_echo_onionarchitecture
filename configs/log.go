package configs

const (
	LogPath     = "/log"
	LogSize     = 10   // MB
	LogBucket   = 5    // 個数
	LogAge      = 60   // days
	LogCompress = true // 圧縮

	AccessLogFormat = "request_id=${id}, time=${time_rfc3339_nano}, method=${method}, host${host}, uri=${uri}, status=${status}, error=${error}, referer=${referer}, remote_ip=${remote_ip}, user_agent=${user_agent}, latency=${latency}, latency_human=${latency_human}, bytes_in=${bytes_in}, bytes_out=${bytes_out}\n"
)

var LogPaths = []string{
	LogPath,
}
