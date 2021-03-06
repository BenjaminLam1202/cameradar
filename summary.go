package cameradar

import (
	"github.com/Ullaakut/disgo/style"
)

// PrintStreams prints information on each stream.
func (s *Scanner) PrintStreams(streams []Stream) {
	success := 0
	if len(streams) == 0 {
		s.term.Infof("%s No streams were found. Please make sure that your target is on an accessible network.\n", style.Failure(style.SymbolCross))
	} else {

		for _, stream := range streams {
			if stream.Available {
				s.term.Infof("%s\tDevice RTSP URL:\t%s\n", style.Success(style.SymbolRightTriangle), style.Link(GetCameraRTSPURL(stream)))
				// s.term.Infof("\tAvailable:\t\t%s\n", style.Success(style.SymbolCheck))
				success++
			} else {
				s.term.Infof("%s\tAdmin panel URL:\t%s URL simply like this but Im totally not sure.\n", style.Failure(style.SymbolCross), style.Link(GetCameraAdminPanelURL(stream)))
				s.term.Infof("\tAvailable:\t\t%s\n", style.Failure(style.SymbolCross))
			}

			if len(stream.Device) > 0 {
				s.term.Infof("\tDevice model:\t\t%s\n\n", stream.Device)
			}

			s.term.Infof("\tIP address:\t\t%s\n", stream.Address)
			s.term.Infof("\tRTSP port:\t\t%d\n", stream.Port)

			s.term.Infoln("\tRTSP routes:")
			if stream.RouteFound {
				for _, route := range stream.Routes {
					s.term.Infoln(style.Success("\t\t\t\t/" + route))
				}
			} else {
				s.term.Infoln(style.Failure("not found"))
			}
			s.term.Infof("%s Streams were found connected. Checking status.\n", style.Success(style.SymbolCheck))
			s.term.Info("\n\n")
		}
	}
	if success > 1 {
		s.term.Infof("%s Successful attack: %s devices were accessed", style.Success(style.SymbolCheck), style.Success(len(streams)))
	} else if success == 1 {
		s.term.Infof("%s Successful attack: %s device was accessed", style.Success(style.SymbolCheck), style.Success("one"))
	}

}
