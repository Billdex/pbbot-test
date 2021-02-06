package pbbot_scheduler

type CmdGroup struct {
	Name         []string
	Handlers     []HandleFunc
	subCmdGroups []*CmdGroup
	scheduler    *Scheduler
}

//func (group *CmdGroup) TrimmedGroup() {
//
//}

func (s *Scheduler) Use(handleFunc ...HandleFunc) {

}

func (s *Scheduler) Bind(handler HandleFunc, keywords ...string) {

}

func (group *CmdGroup) Group(name string, handlers ...HandleFunc) *CmdGroup {
	cmdGroup := &CmdGroup{
		Name:      []string{name},
		Handlers:  group.combineHandlers(handlers),
		scheduler: group.scheduler,
	}
	group.subCmdGroups = append(group.subCmdGroups, cmdGroup)
	return cmdGroup
}

func (group *CmdGroup) Alias(alias ...string) *CmdGroup {
	group.Name = append(group.Name, alias...)
	return group
}

func (group *CmdGroup) SearchHandlerChain(message string) ([]HandleFunc, string) {

}
