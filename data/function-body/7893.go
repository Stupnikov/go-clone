{
	_, err := NewRR("_443._tcp.example.org. IN TLSA 0 0 0 308205e8308204d0a00302010202100411de8f53b462f6a5a861b712ec6b59300d06092a864886f70d01010b05003070310b300906035504061302555331153013060355040a130c446967694365727420496e6331193017060355040b13107777772e64696769636572742e636f6d312f302d06035504031326446967694365727420534841322048696768204173737572616e636520536572766572204341301e170d3134313130363030303030305a170d3135313131333132303030305a3081a5310b3009060355040613025553311330110603550408130a43616c69666f726e6961311430120603550407130b4c6f7320416e67656c6573313c303a060355040a1333496e7465726e657420436f72706f726174696f6e20666f722041737369676e6564204e616d657320616e64204e756d6265727331133011060355040b130a546563686e6f6c6f6779311830160603550403130f7777772e6578616d706c652e6f726730820122300d06092a864886f70d01010105000382010f003082010a02820101009e663f52a3d18cb67cdfed547408a4e47e4036538988da2798da3b6655f7240d693ed1cb3fe6d6ad3a9e657ff6efa86b83b0cad24e5d31ff2bf70ec3b78b213f1b4bf61bdc669cbbc07d67154128ca92a9b3cbb4213a836fb823ddd4d7cc04918314d25f06086fa9970ba17e357cca9b458c27eb71760ab95e3f9bc898ae89050ae4d09ba2f7e4259d9ff1e072a6971b18355a8b9e53670c3d5dbdbd283f93a764e71b3a4140ca0746090c08510e2e21078d7d07844bf9c03865b531a0bf2ee766bc401f6451c5a1e6f6fb5d5c1d6a97a0abe91ae8b02e89241e07353909ccd5b41c46de207c06801e08f20713603827f2ae3e68cf15ef881d7e0608f70742e30203010001a382024630820242301f0603551d230418301680145168ff90af0207753cccd9656462a212b859723b301d0603551d0e04160414b000a7f422e9b1ce216117c4c46e7164c8e60c553081810603551d11047a3078820f7777772e6578616d706c652e6f7267820b6578616d706c652e636f6d820b6578616d706c652e656475820b6578616d706c652e6e6574820b6578616d706c652e6f7267820f7777772e6578616d706c652e636f6d820f7777772e6578616d706c652e656475820f7777772e6578616d706c652e6e6574300e0603551d0f0101ff0404030205a0301d0603551d250416301406082b0601050507030106082b0601050507030230750603551d1f046e306c3034a032a030862e687474703a2f2f63726c332e64696769636572742e636f6d2f736861322d68612d7365727665722d67332e63726c3034a032a030862e687474703a2f2f63726c342e64696769636572742e636f6d2f736861322d68612d7365727665722d67332e63726c30420603551d20043b3039303706096086480186fd6c0101302a302806082b06010505070201161c68747470733a2f2f7777772e64696769636572742e636f6d2f43505330818306082b0601050507010104773075302406082b060105050730018618687474703a2f2f6f6373702e64696769636572742e636f6d304d06082b060105050730028641687474703a2f2f636163657274732e64696769636572742e636f6d2f446967694365727453484132486967684173737572616e636553657276657243412e637274300c0603551d130101ff04023000300d06092a864886f70d01010b050003820101005eac2124dedb3978a86ff3608406acb542d3cb54cb83facd63aec88144d6a1bf15dbf1f215c4a73e241e582365cba9ea50dd306541653b3513af1a0756c1b2720e8d112b34fb67181efad9c4609bdc670fb025fa6e6d42188161b026cf3089a08369c2f3609fc84bcc3479140c1922ede430ca8dbac2b2a3cdacb305ba15dc7361c4c3a5e6daa99cb446cb221b28078a7a944efba70d96f31ac143d959bccd2fd50e30c325ea2624fb6b6dbe9344dbcf133bfbd5b4e892d635dbf31596451672c6b65ba5ac9b3cddea92b35dab1065cae3c8cb6bb450a62ea2f72ea7c6bdc7b65fa09b012392543734083c7687d243f8d0375304d99ccd2e148966a8637a6797")
	if err == nil {
		t.Fatalf("token overflow should return an error")
	}
	t.Logf("err: %s\n", err)
}