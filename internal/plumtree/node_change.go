package plumtree

import (
	. "snow/common"
)

func (s *Server) ApplyJoin() {
	s.PlumTreeBroadcast([]byte{}, NodeJoin)
}

func (s *Server) ApplyLeave() {

	s.PlumTreeBroadcast([]byte{}, NodeLeave)
	//go func() {
	//	time.Sleep(5 * time.Second)
	//	//进行下线操作
	//	stop := struct{}{}
	//	s.StopCh <- stop
	//	s.Close()
	//	s.Member.Clean()
	//	s.Server.IsClosed = true
	//}()

}
