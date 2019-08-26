package hierarchy

var hierarchyTestCases = []struct {
	input    uint64
	expected []uint64
	err      string
}{
	{
		1,
		[]uint64{2, 3, 4, 5, 6, 7},
		"",
	},
	{
		3,
		[]uint64{2, 5, 7},
		"",
	},
	{
		8,
		[]uint64{2, 3, 4, 5, 6, 7},
		"",
	},
	{
		6,
		[]uint64{2, 3, 5, 7},
		"",
	},
	{
		2,
		[]uint64{},
		"",
	},
	{
		7,
		[]uint64{},
		"",
	},
	{
		10,
		[]uint64{},
		"invalid user id",
	},
	// test case for the invalid input scenario, causes a compilation error
	// {
	// 	-10,
	// 	[]uint64{},
	// 	"invalid user id",
	// },
}
