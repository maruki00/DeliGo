package command

import sharedvo "deligo/internal/shared/valueobject"

//  string ID 						= 1 [json_name = "id"];
// 	string Label 				= 2 [json_name = "label"];
// 	float Price 					= 3 [json_name = "price"];
// 	uint32 Qty 				= 4 [json_name = "qty"];

type SaveProductCommand struct {
	ID    sharedvo.ID
	Label string
	Price float32
	Qty   uint32
}

func (_this *SaveProductCommand) Name() string {
	return "SaveProductCommand"
}
