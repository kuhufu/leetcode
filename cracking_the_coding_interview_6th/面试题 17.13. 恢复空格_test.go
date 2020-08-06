package cracking_the_coding_interview_6th

import (
	"testing"
)

func Test_respace(t *testing.T) {
	Run(t, respace, []Test{
		{
			Args{
				[]string{"aaysaayayaasyya", "yyas", "yayysaaayasasssy", "yaasassssssayaassyaayaayaasssasysssaaayysaaasaysyaasaaaaaasayaayayysasaaaa", "aya", "sya", "ysasasy", "syaaaa", "aaaas", "ysa", "a", "aasyaaassyaayaayaasyayaa", "ssaayayyssyaayyysyayaasaaa", "aya", "aaasaay", "aaaa", "ayyyayssaasasysaasaaayassasysaaayaassyysyaysaayyasayaaysyyaasasasaayyasasyaaaasysasy", "aaasa", "ysayssyasyyaaasyaaaayaaaaaaaaassaaa", "aasayaaaayssayyaayaaaaayaaays", "s"},
				"asasayaayaassayyayyyyssyaassasaysaaysaayaaaaysyaaaa",
			},
			Want{7},
		},
		{
			Args{[]string{"looked", "just", "like", "her", "brother"}, "jesslookedjustliketimherbrother"},
			Want{7},
		},
	})
}
