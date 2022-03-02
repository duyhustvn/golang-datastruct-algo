
(dap-register-debug-template
  "Launch Unoptimized Debug Package"
  (list :type "go"
        :request "launch"
        :name "Launch Unoptimized Debug Package"
        :mode "debug"
        :program "${workspaceFolder}/main.go"
        :buildFlags "-gcflags '-N -l'"
        :args nil
        :env nil
        :envFile nil))
