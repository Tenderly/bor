package chains

import (
	"math/big"

	"github.com/tenderly/bor/common"
	"github.com/tenderly/bor/core"
	"github.com/tenderly/bor/params"
)

var mainnetBor = &Chain{
	Hash:      common.HexToHash("0xa9c28ce2141b56c474f1dc504bee9b01eb1bd7d1a507580d5519d4437a97de1b"),
	NetworkId: 137,
	Genesis: &core.Genesis{
		Config: &params.ChainConfig{
			ChainID:             big.NewInt(137),
			HomesteadBlock:      big.NewInt(0),
			DAOForkBlock:        nil,
			DAOForkSupport:      true,
			EIP150Block:         big.NewInt(0),
			EIP155Block:         big.NewInt(0),
			EIP158Block:         big.NewInt(0),
			ByzantiumBlock:      big.NewInt(0),
			ConstantinopleBlock: big.NewInt(0),
			PetersburgBlock:     big.NewInt(0),
			IstanbulBlock:       big.NewInt(3395000),
			MuirGlacierBlock:    big.NewInt(3395000),
			BerlinBlock:         big.NewInt(14750000),
			LondonBlock:         big.NewInt(23850000),
			ShanghaiBlock:       big.NewInt(50523000),
			Bor: &params.BorConfig{
				JaipurBlock: big.NewInt(23850000),
				DelhiBlock:  big.NewInt(38189056),
				IndoreBlock: big.NewInt(44934656),
				StateSyncConfirmationDelay: map[string]uint64{
					"44934656": 128,
				},
				Period: map[string]uint64{
					"0": 2,
				},
				ProducerDelay: map[string]uint64{
					"0":        6,
					"38189056": 4,
				},
				Sprint: map[string]uint64{
					"0":        64,
					"38189056": 16,
				},
				BackupMultiplier: map[string]uint64{
					"0": 2,
				},
				ValidatorContract:     "0x0000000000000000000000000000000000001000",
				StateReceiverContract: "0x0000000000000000000000000000000000001001",
				OverrideStateSyncRecords: map[string]int{
					"14949120": 8,
					"14949184": 0,
					"14953472": 0,
					"14953536": 5,
					"14953600": 0,
					"14953664": 0,
					"14953728": 0,
					"14953792": 0,
					"14953856": 0,
				},
				BurntContract: map[string]string{
					"23850000": "0x70bca57f4579f58670ab2d18ef16e02c17553c38",
					"50523000": "0x7A8ed27F4C30512326878652d20fC85727401854",
				},
				BlockAlloc: map[string]interface{}{
					// write as interface since that is how it is decoded in genesis
					"22156660": map[string]interface{}{
						"0000000000000000000000000000000000001010": map[string]interface{}{
							"balance": "0x0",
							"code":    "0x60806040526004361061019c5760003560e01c806377d32e94116100ec578063acd06cb31161008a578063e306f77911610064578063e306f77914610a7b578063e614d0d614610aa6578063f2fde38b14610ad1578063fc0c546a14610b225761019c565b8063acd06cb31461097a578063b789543c146109cd578063cc79f97b14610a505761019c565b80639025e64c116100c65780639025e64c146107c957806395d89b4114610859578063a9059cbb146108e9578063abceeba21461094f5761019c565b806377d32e94146106315780638da5cb5b146107435780638f32d59b1461079a5761019c565b806347e7ef24116101595780637019d41a116101335780637019d41a1461053357806370a082311461058a578063715018a6146105ef578063771282f6146106065761019c565b806347e7ef2414610410578063485cc9551461046b57806360f96a8f146104dc5761019c565b806306fdde03146101a15780631499c5921461023157806318160ddd1461028257806319d27d9c146102ad5780632e1a7d4d146103b1578063313ce567146103df575b600080fd5b3480156101ad57600080fd5b506101b6610b79565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156101f65780820151818401526020810190506101db565b50505050905090810190601f1680156102235780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561023d57600080fd5b506102806004803603602081101561025457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610bb6565b005b34801561028e57600080fd5b50610297610c24565b6040518082815260200191505060405180910390f35b3480156102b957600080fd5b5061036f600480360360a08110156102d057600080fd5b81019080803590602001906401000000008111156102ed57600080fd5b8201836020820111156102ff57600080fd5b8035906020019184600183028401116401000000008311171561032157600080fd5b9091929391929390803590602001909291908035906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610c3a565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6103dd600480360360208110156103c757600080fd5b8101908080359060200190929190505050610caa565b005b3480156103eb57600080fd5b506103f4610dfc565b604051808260ff1660ff16815260200191505060405180910390f35b34801561041c57600080fd5b506104696004803603604081101561043357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610e05565b005b34801561047757600080fd5b506104da6004803603604081101561048e57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610fc1565b005b3480156104e857600080fd5b506104f1611090565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561053f57600080fd5b506105486110b6565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561059657600080fd5b506105d9600480360360208110156105ad57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506110dc565b6040518082815260200191505060405180910390f35b3480156105fb57600080fd5b506106046110fd565b005b34801561061257600080fd5b5061061b6111cd565b6040518082815260200191505060405180910390f35b34801561063d57600080fd5b506107016004803603604081101561065457600080fd5b81019080803590602001909291908035906020019064010000000081111561067b57600080fd5b82018360208201111561068d57600080fd5b803590602001918460018302840111640100000000831117156106af57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192905050506111d3565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561074f57600080fd5b50610758611358565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156107a657600080fd5b506107af611381565b604051808215151515815260200191505060405180910390f35b3480156107d557600080fd5b506107de6113d8565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561081e578082015181840152602081019050610803565b50505050905090810190601f16801561084b5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561086557600080fd5b5061086e611411565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156108ae578082015181840152602081019050610893565b50505050905090810190601f1680156108db5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610935600480360360408110156108ff57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061144e565b604051808215151515815260200191505060405180910390f35b34801561095b57600080fd5b50610964611474565b6040518082815260200191505060405180910390f35b34801561098657600080fd5b506109b36004803603602081101561099d57600080fd5b8101908080359060200190929190505050611501565b604051808215151515815260200191505060405180910390f35b3480156109d957600080fd5b50610a3a600480360360808110156109f057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019092919080359060200190929190505050611521565b6040518082815260200191505060405180910390f35b348015610a5c57600080fd5b50610a65611541565b6040518082815260200191505060405180910390f35b348015610a8757600080fd5b50610a90611546565b6040518082815260200191505060405180910390f35b348015610ab257600080fd5b50610abb61154c565b6040518082815260200191505060405180910390f35b348015610add57600080fd5b50610b2060048036036020811015610af457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506115d9565b005b348015610b2e57600080fd5b50610b376115f6565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b60606040518060400160405280600b81526020017f4d6174696320546f6b656e000000000000000000000000000000000000000000815250905090565b6040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f44697361626c656420666561747572650000000000000000000000000000000081525060200191505060405180910390fd5b6000601260ff16600a0a6402540be40002905090565b60006040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f44697361626c656420666561747572650000000000000000000000000000000081525060200191505060405180910390fd5b60003390506000610cba826110dc565b9050610cd18360065461161c90919063ffffffff16565b600681905550600083118015610ce657508234145b610d58576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f496e73756666696369656e7420616d6f756e740000000000000000000000000081525060200191505060405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff16600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167febff2602b3f468259e1e99f613fed6691f3a6526effe6ef3e768ba7ae7a36c4f8584610dd4876110dc565b60405180848152602001838152602001828152602001935050505060405180910390a3505050565b60006012905090565b610e0d611381565b610e1657600080fd5b600081118015610e535750600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614155b610ea8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526023815260200180611da76023913960400191505060405180910390fd5b6000610eb3836110dc565b905060008390508073ffffffffffffffffffffffffffffffffffffffff166108fc849081150290604051600060405180830381858888f19350505050158015610f00573d6000803e3d6000fd5b50610f168360065461163c90919063ffffffff16565b6006819055508373ffffffffffffffffffffffffffffffffffffffff16600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f4e2ca0515ed1aef1395f66b5303bb5d6f1bf9d61a353fa53f73f8ac9973fa9f68585610f98896110dc565b60405180848152602001838152602001828152602001935050505060405180910390a350505050565b600760009054906101000a900460ff1615611027576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526023815260200180611d846023913960400191505060405180910390fd5b6001600760006101000a81548160ff02191690831515021790555080600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061108c8261165b565b5050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008173ffffffffffffffffffffffffffffffffffffffff16319050919050565b611105611381565b61110e57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b60065481565b60008060008060418551146111ee5760009350505050611352565b602085015192506040850151915060ff6041860151169050601b8160ff16101561121957601b810190505b601b8160ff16141580156112315750601c8160ff1614155b156112425760009350505050611352565b60018682858560405160008152602001604052604051808581526020018460ff1660ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa15801561129f573d6000803e3d6000fd5b505050602060405103519350600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16141561134e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f4572726f7220696e2065637265636f766572000000000000000000000000000081525060200191505060405180910390fd5b5050505b92915050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614905090565b6040518060400160405280600181526020017f890000000000000000000000000000000000000000000000000000000000000081525081565b60606040518060400160405280600581526020017f4d41544943000000000000000000000000000000000000000000000000000000815250905090565b6000813414611460576000905061146e565b61146b338484611753565b90505b92915050565b6040518060800160405280605b8152602001611e1c605b91396040516020018082805190602001908083835b602083106114c357805182526020820191506020810190506020830392506114a0565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040528051906020012081565b60056020528060005260406000206000915054906101000a900460ff1681565b600061153761153286868686611b10565b611be6565b9050949350505050565b608981565b60015481565b604051806080016040528060528152602001611dca605291396040516020018082805190602001908083835b6020831061159b5780518252602082019150602081019050602083039250611578565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040528051906020012081565b6115e1611381565b6115ea57600080fd5b6115f38161165b565b50565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008282111561162b57600080fd5b600082840390508091505092915050565b60008082840190508381101561165157600080fd5b8091505092915050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561169557600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000803073ffffffffffffffffffffffffffffffffffffffff166370a08231866040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b1580156117d357600080fd5b505afa1580156117e7573d6000803e3d6000fd5b505050506040513d60208110156117fd57600080fd5b8101908080519060200190929190505050905060003073ffffffffffffffffffffffffffffffffffffffff166370a08231866040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561188f57600080fd5b505afa1580156118a3573d6000803e3d6000fd5b505050506040513d60208110156118b957600080fd5b810190808051906020019092919050505090506118d7868686611c30565b8473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff16600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167fe6497e3ee548a3372136af2fcb0696db31fc6cf20260707645068bd3fe97f3c48786863073ffffffffffffffffffffffffffffffffffffffff166370a082318e6040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b1580156119df57600080fd5b505afa1580156119f3573d6000803e3d6000fd5b505050506040513d6020811015611a0957600080fd5b81019080805190602001909291905050503073ffffffffffffffffffffffffffffffffffffffff166370a082318e6040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b158015611a9757600080fd5b505afa158015611aab573d6000803e3d6000fd5b505050506040513d6020811015611ac157600080fd5b8101908080519060200190929190505050604051808681526020018581526020018481526020018381526020018281526020019550505050505060405180910390a46001925050509392505050565b6000806040518060800160405280605b8152602001611e1c605b91396040516020018082805190602001908083835b60208310611b625780518252602082019150602081019050602083039250611b3f565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405160208183030381529060405280519060200120905060405181815273ffffffffffffffffffffffffffffffffffffffff8716602082015285604082015284606082015283608082015260a0812092505081915050949350505050565b60008060015490506040517f190100000000000000000000000000000000000000000000000000000000000081528160028201528360228201526042812092505081915050919050565b3073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611cd2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f63616e27742073656e6420746f204d524332300000000000000000000000000081525060200191505060405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f19350505050158015611d18573d6000803e3d6000fd5b508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a350505056fe54686520636f6e747261637420697320616c726561647920696e697469616c697a6564496e73756666696369656e7420616d6f756e74206f7220696e76616c69642075736572454950373132446f6d61696e28737472696e67206e616d652c737472696e672076657273696f6e2c75696e7432353620636861696e49642c6164647265737320766572696679696e67436f6e747261637429546f6b656e5472616e736665724f726465722861646472657373207370656e6465722c75696e7432353620746f6b656e49644f72416d6f756e742c6279746573333220646174612c75696e743235362065787069726174696f6e29a265627a7a72315820a4a6f71a98ac3fc613c3a8f1e2e11b9eb9b6b39f125f7d9508916c2b8fb02c7164736f6c63430005100032",
						},
					},
					"50523000": map[string]interface{}{
						"0x0000000000000000000000000000000000001001": map[string]interface{}{
							"balance": "0x0",
							"code":    "0x608060405234801561001057600080fd5b506004361061005e576000357c01000000000000000000000000000000000000000000000000000000009004806319494a17146100635780633434735f146100fe5780635407ca6714610148575b600080fd5b6100e46004803603604081101561007957600080fd5b8101908080359060200190929190803590602001906401000000008111156100a057600080fd5b8201836020820111156100b257600080fd5b803590602001918460018302840111640100000000831117156100d457600080fd5b9091929391929390505050610166565b604051808215151515815260200191505060405180910390f35b6101066104d3565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6101506104eb565b6040518082815260200191505060405180910390f35b600073fffffffffffffffffffffffffffffffffffffffe73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461021d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f4e6f742053797374656d2041646465737321000000000000000000000000000081525060200191505060405180910390fd5b606061027461026f85858080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050506104f1565b61051f565b905060006102958260008151811061028857fe5b60200260200101516105fc565b90508060016000540114610311576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f537461746549647320617265206e6f742073657175656e7469616c000000000081525060200191505060405180910390fd5b600080815480929190600101919050555060006103418360018151811061033457fe5b602002602001015161066d565b905060606103628460028151811061035557fe5b6020026020010151610690565b905061036d8261071c565b156104c8576000624c4b409050606084836040516024018083815260200180602001828103825283818151815260200191508051906020019080838360005b838110156103c75780820151818401526020810190506103ac565b50505050905090810190601f1680156103f45780820380516001836020036101000a031916815260200191505b5093505050506040516020818303038152906040527f26c53bea000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050905060008082516020840160008887f19650847f5a22725590b0a51c923940223f7458512164b1113359a735e86e7f27f44791ee88604051808215151515815260200191505060405180910390a250505b505050509392505050565b73fffffffffffffffffffffffffffffffffffffffe81565b60005481565b6104f961099c565b600060208301905060405180604001604052808451815260200182815250915050919050565b606061052a82610735565b61053357600080fd5b600061053e83610783565b905060608160405190808252806020026020018201604052801561057c57816020015b6105696109b6565b8152602001906001900390816105615790505b509050600061058e85602001516107f4565b8560200151019050600080600090505b848110156105ef576105af8361087d565b91506040518060400160405280838152602001848152508482815181106105d257fe5b60200260200101819052508183019250808060010191505061059e565b5082945050505050919050565b600080826000015111801561061657506021826000015111155b61061f57600080fd5b600061062e83602001516107f4565b9050600081846000015103905060008083866020015101905080519150602083101561066157826020036101000a820491505b81945050505050919050565b6000601582600001511461068057600080fd5b610689826105fc565b9050919050565b606060008260000151116106a357600080fd5b60006106b283602001516107f4565b905060008184600001510390506060816040519080825280601f01601f1916602001820160405280156106f45781602001600182028038833980820191505090505b5090506000816020019050610710848760200151018285610935565b81945050505050919050565b600080823b905060008163ffffffff1611915050919050565b6000808260000151141561074c576000905061077e565b60008083602001519050805160001a915060c060ff168260ff1610156107775760009250505061077e565b6001925050505b919050565b6000808260000151141561079a57600090506107ef565b600080905060006107ae84602001516107f4565b84602001510190506000846000015185602001510190505b808210156107e8576107d78261087d565b8201915082806001019350506107c6565b8293505050505b919050565b600080825160001a9050608060ff16811015610814576000915050610878565b60b860ff16811080610839575060c060ff168110158015610838575060f860ff1681105b5b15610848576001915050610878565b60c060ff168110156108685760018060b80360ff16820301915050610878565b60018060f80360ff168203019150505b919050565b6000806000835160001a9050608060ff1681101561089e576001915061092b565b60b860ff168110156108bb576001608060ff16820301915061092a565b60c060ff168110156108eb5760b78103600185019450806020036101000a85510460018201810193505050610929565b60f860ff1681101561090857600160c060ff168203019150610928565b60f78103600185019450806020036101000a855104600182018101935050505b5b5b5b8192505050919050565b600081141561094357610997565b5b602060ff1681106109735782518252602060ff1683019250602060ff1682019150602060ff1681039050610944565b6000600182602060ff16036101000a03905080198451168184511681811785525050505b505050565b604051806040016040528060008152602001600081525090565b60405180604001604052806000815260200160008152509056fea265627a7a723158208f1ea6fcf63d6911ac5dbfe340be1029614581802c6a750e7d6354b32ce6647c64736f6c63430005110032",
						},
					},
				},
			},
		},
		Nonce:      0,
		Timestamp:  1590824836,
		GasLimit:   10000000,
		Difficulty: big.NewInt(1),
		Mixhash:    common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
		Coinbase:   common.HexToAddress("0x0000000000000000000000000000000000000000"),
		Alloc:      readPrealloc("allocs/mainnet.json"),
	},
	Bootnodes: []string{
		"enode://b8f1cc9c5d4403703fbf377116469667d2b1823c0daf16b7250aa576bacf399e42c3930ccfcb02c5df6879565a2b8931335565f0e8d3f8e72385ecf4a4bf160a@3.36.224.80:30303",
		"enode://8729e0c825f3d9cad382555f3e46dcff21af323e89025a0e6312df541f4a9e73abfa562d64906f5e59c51fe6f0501b3e61b07979606c56329c020ed739910759@54.194.245.5:30303",
		"enode://76316d1cb93c8ed407d3332d595233401250d48f8fbb1d9c65bd18c0495eca1b43ec38ee0ea1c257c0abb7d1f25d649d359cdfe5a805842159cfe36c5f66b7e8@52.78.36.216:30303",
		"enode://681ebac58d8dd2d8a6eef15329dfbad0ab960561524cf2dfde40ad646736fe5c244020f20b87e7c1520820bc625cfb487dd71d63a3a3bf0baea2dbb8ec7c79f1@34.240.245.39:30303",
	},
}
