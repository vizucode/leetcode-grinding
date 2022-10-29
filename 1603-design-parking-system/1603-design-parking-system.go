type ParkingSystem struct {
//     menggunakan array untuk menyimpan slot parkir
    ParkingSlot []int
}


func Constructor(big int, medium int, small int) ParkingSystem {
    return ParkingSystem {
        ParkingSlot: []int{big, medium, small},
    }
}


func (this *ParkingSystem) AddCar(carType int) bool {
    result := false
//     car type 1 == big
//     car type 2 == medium
//     car type 3 == small
    
//     kurangi satu jika slot parkir digunakan
    switch carType {
    case 1:
        if this.ParkingSlot[0] > 0 {
            this.ParkingSlot[0] -= 1
            result = true
        }
        break
    case 2:
        if this.ParkingSlot[1] > 0 {
            this.ParkingSlot[1] -= 1
            result = true
        }
        break
    case 3:
        if this.ParkingSlot[2] > 0 {
            this.ParkingSlot[2] -= 1
            result = true
        }
        break
    }
    return result
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */