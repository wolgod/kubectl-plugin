package app

import (
	"fmt"
	"github.com/gosoon/kubectl-plugin/pkg/types"
	"github.com/gosoon/kubectl-plugin/pkg/utils"
	"strconv"

	v1 "k8s.io/api/core/v1"
)

func getNodeAllocatable(allocatable v1.ResourceList) (float64, float64) {
	nodeCPU := float64(0)
	nodeMemory := float64(0)
	for name, value := range allocatable {
		if string(name) == "cpu" {
			//cpu, _ := strconv.ParseFloat(value.String(), 64)
			// MilliValue returns the value of ceil(q * 1000); this could overflow an int64;
			// if that's a concern, call Value() first to verify the number is small enough.
			cpu := float64(value.MilliValue())
			formatted := fmt.Sprintf("%.2f", cpu/1000) // 结果将为字符串 "1.23"
			rounded, err := strconv.ParseFloat(formatted, 64)
			if err != nil {
				rounded = float64(0)
			}
			nodeCPU = rounded
		} else if string(name) == "memory" {
			memory, _ := utils.ConvertMemoryUnit(value.String())
			nodeMemory += memory
		}
	}
	return nodeCPU, nodeMemory
}

func pickNodeCPURequests(node *types.NodeResourceList) string {
	return fmt.Sprintf("%.2f (%v)", node.CPURequests, node.CPURequestsUsage)
}

func pickNodeMemoryRequests(node *types.NodeResourceList) string {
	return fmt.Sprintf("%.1f (%v)", node.MemoryRequests, node.MemoryRequestsUsage)
}

func pickNodeCPULimits(node *types.NodeResourceList) string {
	return fmt.Sprintf("%.2f (%v)", node.CPULimits, node.CPULimitsUsage)
}

func pickNodeMemoryLimits(node *types.NodeResourceList) string {
	return fmt.Sprintf("%.1f (%v)", node.MemoryLimits, node.MemoryLimitsUsage)
}
func IsNodeReady(node v1.Node) bool {

	for _, condition := range node.Status.Conditions {
		if condition.Type == v1.NodeReady {
			if condition.Status == v1.ConditionTrue {
				return true
			}
		}
	}
	return false
}
