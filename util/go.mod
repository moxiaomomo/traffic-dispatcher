module util

go 1.14

require (
    traffic-dispatcher/model v1.0.0
    traffic-dispatcher/proto v1.0.0
)

replace traffic-dispatcher/model => ../model

replace traffic-dispatcher/proto => ../proto