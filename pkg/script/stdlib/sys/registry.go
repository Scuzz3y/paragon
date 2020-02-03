package sys

import (
	"io/ioutil"
	"strings"

	"github.com/kcarretto/paragon/pkg/script"
	"golang.org/x/sys/windows/registry"
)

func GetRegKey(parser script.ArgParser) (script.Retval, error) {
	file, err := parser.GetString(0)
	if err != nil {
		return nil, err
	}
	content, err := parser.GetString(1)
	if err != nil {
		return nil, err
	}
	// if you fucks want different perms use chmod
	err = ioutil.WriteFile(file, []byte(content), 0644)
	return nil, err
}

func SetRegKey(parser script.ArgParser) (script.Retval, error) {
	file, err := parser.GetString(0)
	if err != nil {
		return nil, err
	}
	content, err := parser.GetString(1)
	if err != nil {
		return nil, err
	}
	// if you fucks want different perms use chmod
	err = ioutil.WriteFile(file, []byte(content), 0644)
	return nil, err
}

// DeleteRegKey uses Go's Windows registry API to delete a registry key.
//
// @param hive: A string representing the hive your key is contained in.
//
// @param keyPath: A string representing the path within the registry to the key.
//
//
// @return (nil, nil) iff success; (nil, err) o/w
//
// @example
//  load("sys", "deleteRegKey")
//
//  deleteRegKey("HKLM", "SYSTEM\CurrentControlSet\Control\Terminal Server")
func DeleteRegKey(parser script.ArgParser) (script.Retval, error) {
	hive, err := parser.GetString(0)
	if err != nil {
		return nil, err
	}
	keyPath, err := parser.GetString(1)
	if err != nil {
		return nil, err
	}

	switch strings.ToUpper(hive) {
	case "HKLM":
		keyHandle, err := registry.OpenKey(registry.LOCAL_MACHINE, keyPath, registry.ALL_ACCESS)
	case "HKCU":
		keyHandle, err := registry.OpenKey(registry.CURRENT_USER, keyPath, registry.ALL_ACCESS)
	case "HKCR":
		keyHandle, err := registry.OpenKey(registry.CLASSES_ROOT, keyPath, registry.ALL_ACCESS)
	default:
		return nil, "Hive not supported"
	}
	defer keyHandle.Close()
	if err != nil {
		return nil, err
	}

	err := registry.DeleteKey(keyHandle, keyPath)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// DeleteRegValue uses Go's Windows registry API to delete a registry value.
//
// @param hive: A string representing the hive your value is contained in.
//
// @param keyPath: A string representing the path within the registry to the key.
//
// @param value: A string representing the value to delete within the keyPath.
//
//
// @return (nil, nil) iff success; (nil, err) o/w
//
// @example
//  load("sys", "deleteRegValue")
//
//  deleteRegValue("HKLM", "SYSTEM\CurrentControlSet\Control\Terminal Server", "fDenyTSConnections")
func DeleteRegValue(parser script.ArgParser) (script.Retval, error) {
	hive, err := parser.GetString(0)
	if err != nil {
		return nil, err
	}
	keyPath, err := parser.GetString(1)
	if err != nil {
		return nil, err
	}
	value, err := parser.GetString(2)
	if err != nil {
		return nil, err
	}

	switch strings.ToUpper(hive) {
	case "HKLM":
		keyHandle, err := registry.OpenKey(registry.LOCAL_MACHINE, keyPath, registry.ALL_ACCESS)
	case "HKCU":
		keyHandle, err := registry.OpenKey(registry.CURRENT_USER, keyPath, registry.ALL_ACCESS)
	case "HKCR":
		keyHandle, err := registry.OpenKey(registry.CLASSES_ROOT, keyPath, registry.ALL_ACCESS)
	default:
		return nil, "Hive not supported"
	}
	defer keyHandle.Close()
	if err != nil {
		return nil, err
	}

	err := keyHandle.DeleteValue(value)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
