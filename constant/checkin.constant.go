package constant

const RPKM_CONFIRM = "rpkm-confirm"
const RPKM_DAY_ONE = "rpkm-day-1"
const RPKM_DAY_TWO = "rpkm-day-2"
const FRESHY_NIGHT_CONFIRM = "freshy-night-confirm"
const FRESHY_NIGHT = "freshy-night"

var StaffOnlyCheckin = map[string]struct{}{
	RPKM_DAY_ONE: {},
	RPKM_DAY_TWO: {},
	FRESHY_NIGHT: {},
}
