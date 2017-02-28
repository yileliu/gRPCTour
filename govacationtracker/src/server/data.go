package main

import (
    "govacationtracker/pb"
)

var employees = []pb.Employee{
    pb.Employee{
        Id: 1,
        BadageNumber: 2008,
        FirstName: "Grace",
        LastName: "Decker",
        VacationAccrualRate: 2,
        VacationAccrued: 30
    },
    pb.Employee{
        Id: 2,
        BadageNumber: 2009,
        FirstName: "Amity",
        LastName: "Fuller",
        VacationAccrualRate: 2.3,
        VacationAccrued: 23.4
    },
    pb.Employee{
        Id: 3,
        BadageNumber: 2010,
        FirstName: "Cookie",
        LastName: "Chen",
        VacationAccrualRate: 2.1,
        VacationAccrued: 20
    }
}