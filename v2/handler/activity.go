package handler

type ActivityHandler struct{}

var Activity = new(ActivityHandler)

func (ah ActivityHandler) GetActivityById(c *Ctx) error {
	// aid := I32(c, "aid")
	// if !system.Activity.Exists(aid) {
	// 	return Err(c, enum.ErrActivityNotFound)
	// }
	// act := system.Cache.Activity(aid)

	return Ok(c, nil)
}

func (ah ActivityHandler) GetActivitiesByUserId(c *Ctx) error {
	return Ok(c, nil)
}

func (ah ActivityHandler) GetActivities(c *Ctx) error {
	return Ok(c, nil)
}

func (ah ActivityHandler) Create(c *Ctx) error {
	return Ok(c, nil)
}

func (ah ActivityHandler) End(c *Ctx) error {
	return Ok(c, nil)
}

func (ah ActivityHandler) Apply(c *Ctx) error {
	return Ok(c, nil)
}

func (ah ActivityHandler) Cancel(c *Ctx) error {
	return Ok(c, nil)
}

func (ah ActivityHandler) Remove(c *Ctx) error {
	return Ok(c, nil)
}

func (ah ActivityHandler) Update(c *Ctx) error {
	return Ok(c, nil)
}
