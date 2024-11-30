tell application "utm"
    --listing virtual machines
    set vms to virtual machines
    repeat with vm in vms
        log vm
    end repeat
end tell