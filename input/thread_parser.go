package input

func (input *Input) ParseThread() (success bool, result Thread) {
	parsed := false

	headerParsed, header := input.ParseThreadHeader()
	if !headerParsed {
		return
	}
	result.Name, result.Id, result.IsDaemon, result.Prio, result.OsPrio, result.Tid, result.Nid, result.ThreadState, result.ConditionAddress = header.Name, header.Id, header.IsDaemon, header.Prio, header.OsPrio, header.Tid, header.Nid, header.ThreadState, header.ConditionAddress

	parsed, result.JavaState, result.JavaStateDetail = input.ParseThreadState()
	if !parsed {
		return
	}
	for {
		lineParsed, line := input.parseStacktraceLine()
		if lineParsed {
			result.Stacktrace = append(result.Stacktrace, line)
		} else {
			break
		}
	}

	if !input.MatchWord("\n") {
		return
	}
	success = true
	return
}

func (input *Input) parseStacktraceLine() (success bool, line StacktraceLine) {
	input.Mark()
	success, line = input.ParseWaitLine()
	if success {
		input.Commit()
		return
	}
	input.Rollback()

	input.Mark()
	success, line = input.ParseBlockedLine()
	if success {
		input.Commit()
		return
	}
	input.Rollback()

	input.Mark()
	success, line = input.ParseLockedLine()
	if success {
		input.Commit()
		return
	}
	input.Rollback()

	input.Mark()
	success, line = input.ParseThreadPosition()
	if success {
		input.Commit()
		return
	}
	input.Rollback()

	input.Mark()
	success, line = input.ParseParkedLine()
	if success {
		input.Commit()
		return
	}
	input.Rollback()
	return
}