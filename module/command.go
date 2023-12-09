package module

import (
    "context"
    "errors"
    "fmt"
    "github.com/filecoin-project/go-jsonrpc"
)

var (
    ErrCommandNotFound       = errors.New("command not found")
    ErrInvalidArgumentLength = errors.New("invalid argument length")
    ErrInternalServerError   = errors.New("internal server error")
    ErrCommandExecError      = errors.New("command exec error")
)

type Result any
type Parameters []string
type Arguments map[string]any
type Executor func(m *Module, args Arguments) (Result, error)

type Command struct {
    Name   string
    Params Parameters
    Exec   Executor
}

type InvokeFunc func(command string, args Arguments) (Result, error)

type ModuleInvokeHandler struct {
    ModuleStore *Store
}

func (h *ModuleInvokeHandler) ModuleInvoke(command string, args Arguments) (Result, error) {
    isFound, mod, cmd := h.ModuleStore.HasRegisteredCommand(command)

    if !isFound {
        return nil, ErrCommandNotFound
    }
    if len(args) != len(cmd.Params) {
        return nil, ErrInvalidArgumentLength
    }

    var client struct {
        HostInvoke InvokeFunc
    }

    closer, err := jsonrpc.NewClient(context.Background(), fmt.Sprintf("http://localhost:%v", mod.RPCPort), "HostInvokeHandler", &client, nil)
    if err != nil {
        return nil, fmt.Errorf("%e: %w", ErrInternalServerError, err)
    }
    defer closer()

    result, err := client.HostInvoke(command, args)
    if err != nil {
        return nil, fmt.Errorf("%e: %w", ErrCommandExecError, err)
    }

    return result, nil
}
