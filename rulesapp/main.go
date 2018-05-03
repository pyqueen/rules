package main

import (
	"fmt"

	"github.com/TIBCOSoftware/bego/common/model"
	"github.com/TIBCOSoftware/bego/ruleapi"
)

func main() {

	fmt.Println("** Welcome to BEGo **")

	//Create Rule, define conditiond and set action callback
	rule := ruleapi.NewRule("* Ensure n1.name is Bob and n2.name matches n1.name ie Bob in this case *")
	fmt.Printf("Rule added: [%s]\n", rule.GetName())
	rule.AddCondition("c1", []model.StreamSource{"n1"}, checkForBob)          // check for name "Bob" in n1
	rule.AddCondition("c2", []model.StreamSource{"n1", "n2"}, checkSameNames) // match the "name" field in both tuples
	//in effect, fire the rule when name field in both tuples in "Bob"
	rule.SetActionFn(myActionFn)

	//Create a RuleSession and add the above Rule
	ruleSession := ruleapi.NewRuleSession()
	ruleSession.AddRule(rule)

	//Now assert a few facts and see if the Rule Action callback fires.
	fmt.Println("Asserting n1 tuple with name=Bob")
	streamTuple1 := model.NewStreamTuple("n1")
	streamTuple1.SetString("name", "Bob")
	ruleSession.Assert(streamTuple1)

	fmt.Println("Asserting n1 tuple with name=Fred")
	streamTuple2 := model.NewStreamTuple("n1")
	streamTuple2.SetString("name", "Fred")
	ruleSession.Assert(streamTuple2)

	fmt.Println("Asserting n2 tuple with name=Fred")
	streamTuple3 := model.NewStreamTuple("n2")
	streamTuple3.SetString("name", "Fred")
	ruleSession.Assert(streamTuple3)

	fmt.Println("Asserting n2 tuple with name=Bob")
	streamTuple4 := model.NewStreamTuple("n2")
	streamTuple4.SetString("name", "Bob")
	ruleSession.Assert(streamTuple4)

	//Retract them
	ruleSession.Retract(streamTuple1)
	ruleSession.Retract(streamTuple2)
	ruleSession.Retract(streamTuple3)
	ruleSession.Retract(streamTuple4)

	//You may delete the rule
	ruleSession.DeleteRule(rule.GetName())

}

func checkForBob(ruleName string, condName string, tuples map[model.StreamSource]model.StreamTuple) bool {
	//This conditions filters on name="Bob"
	streamTuple := tuples["n1"]
	if streamTuple == nil {
		fmt.Println("Should not get a nil tuple in FilterCondition! This is an error")
		return false
	}
	name := streamTuple.GetString("name")
	return name == "Bob"
}

func checkSameNames(ruleName string, condName string, tuples map[model.StreamSource]model.StreamTuple) bool {
	// fmt.Printf("Condition [%s] of Rule [%s] has [%d] tuples\n", condName, ruleName, len(tuples))
	streamTuple1 := tuples["n1"]
	streamTuple2 := tuples["n2"]
	if streamTuple1 == nil || streamTuple2 == nil {
		fmt.Println("Should not get nil tuples here in JoinCondition! This is an error")
		return false
	}
	name1 := streamTuple1.GetString("name")
	name2 := streamTuple2.GetString("name")
	return name1 == name2
}

func myActionFn(ruleName string, tuples map[model.StreamSource]model.StreamTuple) {
	fmt.Printf("Rule fired: [%s]\n", ruleName)
	streamTuple1 := tuples["n1"]
	streamTuple2 := tuples["n2"]
	if streamTuple1 == nil || streamTuple2 == nil {
		fmt.Println("Should not get nil tuples here in Action! This is an error")
	}
	name1 := streamTuple1.GetString("name")
	name2 := streamTuple2.GetString("name")
	fmt.Printf("n1.name = [%s], n2.name = [%s]\n", name1, name2)
}