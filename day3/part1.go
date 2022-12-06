package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	input := "LHLRlCCvCLVgHPfCHtVjBGrBDNzWFBsBGBfscGsD\nnQwbnwwpbrJBrNWB\nhmnSdSdQpTpdnlPdvddPNglLjH\nRZhwpDsNqVmQClwl\nTLJfLTPqcvTrvvLMLMlVzzvVVQQtmQCmtzmV\nMJjccdfTMcbqjNSRSZsSDZ\nLLrNNqCTCwLTttwcNctqFGmRBSBjzjbSzbBbjNbzjB\nGnhhZQPDGdldgQmQSjpzjzQssb\ngDJZPMnPnhlhJWhZntLCLcTqVMLrGVtMfM\nrrBgDBGnVnffDnfQQqngJhhSRQvhhCRRRSZbRpRzwQ\nNtLmcHPHMHHssFJphZpbhwpNRbbC\nLJPHlmdJDgrrqrnl\nnJhrcNnfrFwNhPdMQSgZSCMjQn\nLjqGWsGWllRRlHVsqGGWsZZSSHQgPmHZZSPvdPCmvQ\nzqqVTWjqBsTJprNbppFb\nzSMgWzlgFSWFcGZlCZGlrrTc\nspnQHdQmHddNmpWrpWcChccTWc\nBsRsnmBQdNWsvRPzbzbLzDVSPSbVLM\nlDfbffptlrJZTBJHjjBWjT\nLcwwgQLgzvztwtMQGCMVCHWmnmjWnGhFHnZjmZhjhT\nsCqtzsswCgccbSqrDSqbNNfN\nsnnnjwRRwGSSnVmhhVMhGFbgQgbzFFPPgQQmPbbgQd\nqCrccvcDDcvqDZlCcrcfQNQFdsbgWzFfQddQQPgQ\nZcrvrBqBTCZnBBswjwpSRs\nqSczBfBcjMZMfctsmsGmFJsmQQcQCr\nwPhTLNVNGLNdGHPHwlQsnrnmnrQvHFFHQn\ndLdwbNLRdgGbgTjZfDbqDWjftSzW\nrZwlrtRtNtlHqVBtdqQgdq\nfbwzpPwbhJzpwfTSHgdgqcJVcBjHvHdJ\nLLPbhzPpTTbTshfGhPwSFWnNmMrrZZmNmZDmWNCCZs\ntMHgMWMQWgFJTHsWMvJrVdlmvlSvdvlpvG\nRNfZZfRttBtdlZlmmmplSS\nDzzNDDRwnwwbLnMFtsMntQFM\nqHqBMNqgMwHMbnGStHSbndnt\nPwWZPjpfsDsDsPfPfjdbSvbWhdFSbFGSWFtr\nwfpjpJcfVsspzZRRszDpwcRggLTQQBLqNqcqcggLCgNmlq\nTmmFjtvFdDGjdFFJjFRDLNLHGBVcqgLcLgVBLqgV\nWbWSPSwQCWrWQSrCbwNVlLlBZLBvLlvZVqPl\nSfwbhrwQQbbMwCwWCrbwJRvptJfjDTDRvzTttRjp\njzqSMszqsbVVMVMgWhWCgMrpGgpB\nwrwLcFQmPlFFlwLZmFGTfPvWGfPvhWWWvGgf\nwRZtHFHmzNDHSqrs\nNprlCgrrnrNCjplSCtljpFrFZLzzgwmVgBzBZZPwzBPQBwVL\nHsDMvHTDfsfQZfZzmPWL\nJcPsDGTqcTqSdrSCtnCt\ncFcmfmJFtFmtlTNtLlCWTT\nQPQzjRRsVsQqBqwlTlNBpLNSWDpN\nQblgzRPgbgQsVvgPVQhgQqjvMnJfccnZddcGGfrFJMFGGF\nPWbWmFFnPFPWbDVVmmDHDFGdGhTQdLdnTZQZZcGSGGdQ\nClzjNBlBJvlsBdcPLZdLPQjLQZ\nvJMBpBzzzfNCCzCffJlzgMWDWwPDtVtmVHPMwVHD\nbJjWzWFlTMjjSNBrRcBrZR\nmwnwqPwnGQPCqmJmPQJPCVNcRZBRRrrNrmrcVpSrRf\nPGvQQGPqvhWFWlJbDv\nPNPrdmPGRJlZCrCJlGQzjRFLpFRppjgppgcj\nDwfVnssbVnSWShDwsnnhBLFjFgjFBzDBjHDLpHDj\nvSsMgbsTfTwwfMffnTvgNNCmrPJtCNrJrCrrtvGm\ncRnRplCzccVcrwcnppVVzRCNhfhgChNJfPgHJdHDNtNCtP\nWFmbLMZdLBqfJNbPTfttDD\nBdMWdQsGsSsrpzrswr\nllhhZzSLqlzwRrffzwzT\nGvBbNjHbjjTGGHHFcsFvfRrtJQPvtRfwwfrPJD\nHGbcbTTjHFNpppmLnSdplWqZ\nFhwFbPwsvtRgVCgvMT\nHJVHdHBWdBQSSSQnqSQLqZHtcCctppgBtRrgtMCgTprRMM\nQNzZLVSLLLDGPPzPmbFs\nVdTHmWCVZDTPBBWBQBFQQg\nMzjMjzCjJsbJhhPz\ncrfGGLwwLGrtrvCtdTmdDH\nwRLvLmGQLwFPBRmnLCLmGQTzNNqVNZMMVzzQbzVbNZpMVb\njsgJWjdHghsglHtWsjSfHzVNqzfpCCzqDpzDrVVrrD\njJtWWsWhtSHsSSgchthHcjHCvcFCRvwvvTTFBGvBmmGLnL\nLpjWLNqWpwRWMqLRGjwJlStgbtrVgHFrGllDDrVH\nSQmmTcZZvSZBTmTSzhPTddbVDhHllgFCDHtrDHgDVVDr\nPzTTvznBncnfTmTTQcPdmzzMqWNfqwRLpWJsNfwLJjsSwJ\nlsGdGwBsflGrfsHvHwQLdFrmPhDhCFhhjWCVmhDzmbmPhC\nqtMSNNZZMpcnVzmVbCqjWjzB\nZZcgRJpBtTMNnntncwgQQdfGHHHlsQffLH\njBBtjjqfnwStBSrVVFwSVVvvWzHmcWvWbvPmPbWrbMRM\nGTdNDlpJhlCvPbHgcDmgDH\nJZQdQhNldLdTpGdJGdNCpLZdSBnnBFfHHswqqjffZsqsFjns\nphJhDPQLDSJvpHhvDJhfrFQVRrnsslrgwrVrrRjg\nmWNWqZWWZBMdCGMNCdWmWCNCsnVFTRsVnZFlsrlFwFgVsgjr\ndWdjttGmNCBchJfhHvhPtvJt\nCSFSFdfCzJhtSCHQFjQHQWFHRNHG\nwnbrgZnwZgDLsLbwsLrsrNWQNjPZHHvPPQHNqHHvqB\nLbmTLDsgggQmzmCCppdtSJtM\nSzSSchCdZgHbwHSZ\nGsMMmslnsfmNGNNNVVtZWdwbqQbpgWjjgWZjQm\nDvlMtflGGVGthhzdvLvhrTcc\ndvfVNqHlQfGRcjDczlCDnC\nPsPsStLprtTTFSTLmhSVSFSsCzRRjzDnMJDCMWWDjMnMnpjW\nhBFPhSBFttBhStLwmsPTtPsFHZNGfQgdgdZbdqZwdgNVvwvH\nrhjcChdgjdCrjLjLLSLmLFMmFtNnnbQNNNPMbbmtQF\nlwWRZDlsWzrbbQpN\nrTRqsqDRRRsDRVrqDgBCdCVShjCBHchjdh\nPsspltlPsmTsmbmfTPSTTCGjhJJjCnpqJJNhhJwNJh\nBdrLVvgQLQVLHRZGnqRhRNdwNC\nWrFVHDgDQFHVLVVDFQMLltmPsssPztwsPTzsWcmc\nCBvvSzFGSGGWfFZpcHqjvjcvcqccJq\nbRQwgbbhrRhdwmQbWtdjnJHqVJccJqgHVJHJcl\nRRmbPrNRRPLLtmbQmbNwddCTSTzBSDBDSMFTBSSPZFFW\ncppsSgNrSgwrGRdHRrwd\nLzqqmCLCLWQvCzmzZwHnZZHSwvwnlDlS\nhzFqFLLLFtSNVsFF\nzpZcZZZdppzDLWDtJGgfGbTGGJTGcc\nqhvNSClCShRrRBBWTQfgBFbgtfgg\nHlHqvjqmCvCvlSSHvVdsVDzjpVjMWdwLpP\nqLdsfNsTHQwnSNSBNS\ngFhWzrhfbmlpmZhJWrFSvRMnwwvvpBpBSpQBMv\nrgWZrbmlmbzFfglgWzGggFJLccVPqLPqtPLGcPsPHcPLTd\njTTWRCCbwJJNTHrffqNnzh\nDZVmDpgGBVdcMZnqfhlNHQMlNNhl\nsDcpsDZBcmgdssZcnmSWCPWRSRwJWwvFLvRwWj\nhWwhgQlQQgjPhFChZVdbcJ\nzHsrMPNMtDDTmbcJbccmDb\nznzPzrtHtHtMzqHHrsSSLwljqfgwGggjlQjQQgQBGj\nJpnRtqlJsqDJJBBNNmQmgdmRNGGmvv\nhTCbTwMCwVhTWdmDgDvjWD\nMDhhMSwZCbbLVhbLcDSwCwZtZznlzlnqBPBpHPPlBlHJ\nCtvnvqNNDchrhFVpwftmgQgpQfwS\nMbdqPWGjBjMBbwlfVgdmfSfJJS\njWzbzjWWjWMMbRbMsjzBhChrNHcNqHcrNnhssnsc\nLlLJSWgWllSShRmRlBLJSVBzpTHzTTJcpTHzpTTcPpGpTr\nMfMqnvbvDfbFFZDfFNjsGrRppHpZGGcGrcPprz\nwfNFvtwMvbnntftjfNtVmlgmStBmlBWdQQRg\nGpFRRPGWqzHwdqpzqbjjgfZptBBVMSjSfBZc\nTClllrnsJvDMBgcjfmtssB\nNlhNNchrNJlLvClNDrzbGRqwqqqwqPPFLGdd\nqFmVtvmmVvzzFtzzGzzMNNMSSTjNJlStjSfNgf\nsWrPBCnCTMsTJfSM\nQTLbpnRpRppnRQdRzRZqGzFFVVRz\nWGGPjFvMVNjlcQJr\nbslfldbgtpSmwmSNHQhLJhcwLcQrQV\ngTltCsSsssPFnDzWPTMz\nhhRRhQgGrHjhRsrgqznbzncZjVVJVjncjd\nDSFfNTBFSDmMSTDlFbBBdccCdJJZCbZCbW\nDDLmdSmTvQQgsgvGHH\ndNqNgNvFnvdZHFWnZWNBTQlPTppPGlCTpBQppq\nLJrtLrsLjsGvTCTpQP\nmtLhjVjMhhmVMvtJmLfhFWHnngbRRdZHngnZWZ\nNzdVNzqqCtCHMMZBCGBW\npsjllRjFpjpbjspFmWmWnLBmMMQMmHbm\nDhsTPDRTDHpsvRjdNtzJJJdhqwcdqc\nVbhRbZgRHMFhQpHd\nfvlqPzmzJJqJSPsWmPTNddNFHbNFGHNTHSbc\nCqzlfqrCnbrBZjBr\nSNSrDZFHnTqFsFddsCmsMC\nctVthlGjfhGljcCJmcqMCqcqBB\nVtjvtjhhPPtWqVPLjvqjLVNRppRTvNSnbnZRZTHRnpTD\nfzsBSsNBMNMszNGGJvgjjPggzjdFPgpJ\nbmrVVVrmRrrvRmwvqlbHTDgwdLFjQPJPFddwJPFggj\nHHHrZqhqbTMcZBCZfvcN\ntzsJsnsmBzlVqjssZZrg\nMQZHfNCffpMfpGSpPvpfCGCTTVwFTlwrggjSgqFjVjTwwg\nNfMGGGPZpvLDvLCGGfQHMpZRDRWchRRtBBzmJnmzmnBznb\nBSRBjtNjZrsjRjjNsVBjbrMwCgGCCwCdHrlcdccGcH\nDJTTJLpFnFLdJJqPLTWqLTpwHzCGccCzvvcHwMvWzggCMC\nPmpTTdndmmLqfLTTLDqJVBssbbStVmjjSsZNsBZm\nddCnZvCDSgghFhbbmFVQ\nJzlMcJTMMPPfJJfsMsjWlHVhLbQVlFWmHbbb\nBPwwsPfsqszfFqppwTsqzpntDSnCBnDRZrSZdnDdtvZD\nSllzzPplWldwLGlzbtPZZjVScnnNSjnNsqNqsc\nBrCfFDJFDHBhJCChQFhCCBDcTnNVpZZcNcvQTcjvvcTcZV\nRCmFHJDJhCmBmRCgDCFRpmGbWWbLPlbMWzzGttzgLMbt\nvGTfsZnfvfzTjsnfzTJlwqQjwmCqqMFjFFQMlq\nLHtHRVLRLNtWcmVbRrPbRcwgQwgwMwClwClwrgFpwqpw\nVDPtbVLBmLbLbDDNnnnzJJfJfBfvSGBn\nlpPCRVVQppzHlZgzglgF\nrtfttLdLdscmGtzngPHHFHFH\nLPLLhfhbTDLmPdcrcWdTcDSjjRqwwbqpRwNBNpBwjQwR\ndWQfCJrwvQCfFqNwRbbzVbVVLGTR\nZpZshPMzBjGjtVMN\nphpSlSlDlcZpcZPrdHCFzFzFWFDWQH\ntfMMZhjLlChsdsds\nPHQRMHRwpRPBMvWvPRBpPdWDslGrbscTlTGcCsGddG\nFqPSvHPHPBzQRBBwwRJfVtgjzntMntjJLMtJ\nVBwJvwVwNbVRdPwMgWggGMgH\njDhqflDDhrqshNhdgPGHphLg\nFltsrtcFrclrNqDfqmzQJQQRBzBCvCFvBR\nRZsSSJDJZLDWnGDMLD\nClbnlfmpNtmgbtmMqWdjNGjLQjLqLj\ngcblTlVCnVmcPrvRPFRrZs\nmbJcScmbDWLmSBzwjPRTfjmmRhpl\ntFFFtGttdClHVMCHFMMwTpwNjGpzPpNRRzzpPN\ngCCdvZCVHsFvJnnDSglSglDn\ndSndnRRvVSpLSphfqvTgWqrzqvvw\nPBFQbQbDhGfjTTFzqG\ntPmCJMtDDNcMVdhhVc\nQVRVHCQRmdTRqrZFCWrLZNZFbb\nncncsPnhslBRSSSbFhtbZDLMbLbtLb\nlvflPcfPSsPzlJlPlcPBfHJQwRwHmqdVpRGGGdQmww\nGTCGMCcGdgRnnbbbMLwmMz\nDQFZzllWDDLwDJLnJpnp\nqBVrNNlZfFNlWlWqfRzhgvhCqHRdCGGSvc\nZnMnGbLZfJcBcLTgWF\ndHJjdzqssHHNJwCHpHtDccvtBTtvccWWrrTWSg\nCmlqCNzCzHlmdsqzNzRhMhZRbZMPmRfRJQPb\nLsLLrFLcFjrtmZhhmhHGhJGGhH\nffvbsbWpSBSSCCQbsSBSwwJHHvhZHHGdGVGlMlTVdZlT\npSzWNPSfCwWNPBfsFqtFLtsqRzqFgj\nhwwpvjVppGpwWGLrcPjrbrrdbjdL\nFBqFFMFHHsHssNHtslqtFmldnnLrPhMnccrnMzZnbLgPLz\nmJBSstlJQpwGSVhC\ncgJDVWsrWggpcHhMzwwPnQMWMm\nSZGBjdBqBBjGjjqNGfGNNHPnRSQFzhnwmnQzQnPHMR\nbjjCZddjZbZBCtLhCZhftrgJglcTlJgvTllJvDVDLg\nQpRJpCFdpqTQcqSTBBGBZVjZjVjFvwVB\nnnnWfnHhPDlDnlLwGjBBbVVZGbCGbP\nLhhLMLMrMWMrCprTqpJQpz\ndqGGZJdZbTTMFFTGJFFbMdnCHSdWcmNmcCdWSggfSW\nQsjjtrzLrQwDPjrQLrCfSSnmCmHWlCgmzlNl\nPQpQPjQPsBstLBPttDrjBwTZMFMZFvJFJhMhnMJqpTJJ\nJMLrSvHJdJvvJfrHMJRfWzWDFPwCcWqRRRcq\nljZsZTmmtTBlpTlTjQZCtNFPVqDRwWwWWVPcDRVpFRDz\nmBgQgNNTNvrvJSvgCb\nDWbWtzWDfDffbsMbZMffDDLncnnCJmLVsJJJnhgcngLs\nTjgNGSBRTRTQrFRjFGBVLwLnnNncCLLCCcmhPC\nGgGjvgddvvWqqfdZftWH\nzMmsQlMfQQMhjsmjfsmHlhncRRZnRRRJRvZWWnhccdRC\nBptFtDSSrTrpgtgqqgtZtvVVdVvccVnJdVnG\nqpgPqBLDNTgqBrSLpDBMJfjzmbJMHjLMfjslzH\nDPgLgPhfNDRqhDFDsBTtrrrdbbztCbtf\nMjGSScGVGSlJjbbrtTvdzsTq\nJJwJGWMZwMlWnFFgqNQFpF\nWRGDHmGqWHlrmtVVVRVqpNZvZvvvTNPMPjbPdM\nBwhBwsnzwhzSfCfswFvpvTzTdpMpjvPMZNTb\nLFFQgnbfChSFBhFnftRRLrttDmmRJHtlGH\nMhqhRHmDdRlRlGnfZbJVsNNZDnNb\nQwvzgtwvFpmjwzLjFLJZrsZbPfPZbsVpfPsb\ngvjTzBLztLTjwFjtgLTgtzwdWRqdqRTMSWcWTmWlqdhHHm\nZfzzfmhdpNLNBDDsFfQVCDggfV\nHPFjljSnHrqVDgtgQgQMqC\nrGnSJHvjSwGzwFhGZG\nHqmHRDprrNTZTMbh\nCJvzQRQVQCgNzZbzgMNd\nvPCvFPcfQFlSJBcfRcPHmDGqWGDqpGtjjtGGHl\nwcfpJVHfJBffBBGWRprNRWWWNdhv\nDzMzMPqjDnjgCMZPZjzjCjChGdvvbhNdSvrhNWSNqWRRdS\njjCTtnMTDsMBtLRQwwQlFV\nJqGnVqCTpDVCTnNLgmPzdgjcGmRg\nHrSBJSHblsJthsBBSBhMsrzmdNRccjLzgcLmjgPPjlLL\nHSwSttbswWJrbrSWppvVqvvnQVppQQ\nJDCHssRTTwcRJDcnCDzRHsHNPZGBtPzFPSPttZSZGBqPBZ\nvWhLmTlfrhFqGWSNQNqF\nvpMhhpfvmTfhvbLhhgvmgvvlCCJCCMnDnDCnsjRMVDMDswCC\nZgjdlmlmmlJgHJlbZrSDrnrMrmLLDFprmp\nTvqdTtdctvvDrGGSDn\nCWPBhtWqPPwcdVwlNJfVVNNbbb\nvgmrrwlPPrwPBPtmvFcMMrsMSJHscJcMSHDH\nTWdLnZjCLGLMQLHBLS\njqVTTZqjdTVjNFNPqBvgvBNl\njmcgMzsmjmfvJwFpFfRWZRWp\ndrdSldTmCmTDCNCtbRRWqRwtttFZpZWZqw\nNLCVLNLTbbTrQNQDvnzschgghnmnHQcH\nsRVhVQDVDQRRMQhsqtRRNzqzbNzRqNGp\nWdjCLHLjdFnCCnjnFnLHHmPmNJbztJJpprBpbGzbpJbqGWtB\nCCFFCnFjdnjCTHmCLTLLCnFnQVhQQVDMhQQVgZZNVVDDsTVh\ncGLzZgfzcNNzzRZvjvRmVDmmqCCDSdVVChVnDf\nQstsPlWHQlWhMMtpsbtpQtlpDqBVqPSCSVnTTqmVdDVDVdBV\nstptQFJrHptlQsFJMHtHhFbvNwvjLGvvNgzgjcwczzJcNJ\nSQHCrCFPJZcnWrqn\nvfJvJjfGGDggqZGcWD\npjLpRwzhRFtHdMQJ\nHNSHNDvnvdffDNfqdZfdStcFFMGmmrRBcmFcrMrqrWFB\nVblwzwhwPLlCGGzgzhmFjbjFrMFrjFBmbcrM\nhlVPTCCQCCzVlPzhGgPVJCpHtdtSpQSZNNfnZdSSdnSD\ncBVmfwqwmWggTRmQzTQl\nCDnnHjSDPLCSCLHLHCHNCDFgwJljFQRRhlglbzJFQQhl\nGptHLtNGHtCHSnCMtGWdcsqqWwMcqvdsVwfd\nHsMFNRNWnbWfZLzWzQ\nPqrpjqNdjjhPcdbpvmfzfbffbzmv\nGjhhcPjPccVrqcPCldCjssHNnTnNttRwwVMTHMnV\njjCcBswcfnwgpPFPwFFGSFSwFb\nVvmVhVvRRQqttRtQDLzLhqRlrBFMWrGPBSMSZSMFqBWZZP\nmvQmLRvJtBJVnTcJCjsjsdgp\nqgqvPbdMDMMPwpbLFGwtNlNF\nTTQmdJTnSllFGtJNtw\nCSSHmQHfmVcHjQmSvPBdBDDWMVhMhRMB\nWCvQNZdhCDnnPfQPfTzjHcppsHjpsSjHNS\nmFMgMBlMmBqHjjBfTjHzBc\ngrJbrqfqMVFJrlJrtCtvhPPQdnvnnvnwdW\nFdQQJRdfSSfrJsRZfFsRSvtDBmDHGtGqbgvnmbDnvDGq\nlcMzjCPTVtMqgWWGqn\nVpLjcVPczhzznPLcPhTwFQFRZNfFNrNZFpZNsZJR\nVWgJhmdDdJDdVPggPSTSTWvvfRzfFfFbNb\njCQtnpGQrHMctnpzRbFfgSwHvgwfwv\nntcMcqLMcQccQLgjBPLdhZDVlJPVdDLJ\nRnPnwtqHnJthjLMcWWncMn\nmsdCrCdNpBBsCrlNTpNBDNGzcLchQjFQzccQLQpzLzWtFS\nsTbdTBNrTCTTBCNBlVbwgVPJHtPgPvqgff\nQmBsmpmcZQNqPqVnPFVpGh\ngDDMDLMJgHfJwJMzfTfdGLhChPtvnGRPRRLFGPGv\nDTlzgwfDrrrMWlncbscNnlSW\ntBwvGHFttrFrvRgRhCmCmwQmMg\nJWbNJZjzfbVjWjBhqfmSnhqCqgnQ\nZZJJJbclzJcsTPdvTTPTBFtHDF\nLszmFTFpTmszLrpqSmFpzcvQjtQjvLJgJtcBjgtJjj\nVHNwwNCVCChddfwHlWdnlnGRQPcQjRvMWBJJtMMWcvPJMM\nnGHNVHhnfnHDNhCfdhNNlwHvmpDrZDmpzmbZSZFsmmbqrrsz"
	s := 0

	for _, i := range strings.Split(input, "\n") {
		m := map[rune]int{}
		for _, j := range i[:len(i)/2] {
			m[j]++
		}
		for _, j := range i[len(i)/2:] {
			if m[j] > 0 {
				if unicode.IsUpper(j) {
					s += int(j-'A') + 27
				} else {
					s += int(j-'a') + 1
				}
				break
			}
		}
	}
	fmt.Println(s)
}
