{
entry:
  %"$ret62" = alloca { i8*, i64, i64 }, align 8
  call void @llvm.dbg.value(metadata %TabView.0* %r.19, metadata !933, metadata !DIExpression()), !dbg !934
  %"$ret62.0.sroa_cast16" = bitcast { i8*, i64, i64 }* %"$ret62" to i8*
  call void @llvm.lifetime.start.p0i8(i64 24, i8* nonnull %"$ret62.0.sroa_cast16")
  %"$ret6219" = bitcast { i8*, i64, i64 }* %"$ret62" to i8*
  call void @llvm.memcpy.p0i8.p0i8.i64(i8* nonnull align 8 %"$ret6219", i8* align 16 bitcast ({ i8*, i64, i64 }* @command_line_arguments.fileDescriptor3 to i8*), i64 24, i1 false)
  %call.19 = call i8* @runtime.newobject(i8* nest undef, %_type.0* getelementptr inbounds (%ArrayType.0, %ArrayType.0* @type...61x.7int, i64 0, i32 0)), !dbg !935
  %0 = bitcast i8* %call.19 to i64*, !dbg !935
  store i64 1, i64* %0, align 8, !dbg !935
  call void @llvm.dbg.value(metadata i8* %call.19, metadata !936, metadata !DIExpression(DW_OP_LLVM_fragment, 0, 64)), !dbg !937
  call void @llvm.dbg.value(metadata i64 1, metadata !936, metadata !DIExpression(DW_OP_LLVM_fragment, 64, 64)), !dbg !937
  call void @llvm.dbg.value(metadata i64 1, metadata !936, metadata !DIExpression(DW_OP_LLVM_fragment, 128, 64)), !dbg !937
  %sret.formal.142021 = bitcast { { i8*, i64, i64 }, %IPST.0 }* %sret.formal.14 to i8*
  call void @llvm.memcpy.p0i8.p0i8.i64(i8* nonnull align 8 %sret.formal.142021, i8* nonnull align 8 %"$ret62.0.sroa_cast16", i64 24, i1 false), !dbg !938
  %tmp.27.sroa.2.0.cast.1108.sroa_idx13 = getelementptr inbounds { { i8*, i64, i64 }, %IPST.0 }, { { i8*, i64, i64 }, %IPST.0 }* %sret.formal.14, i64 0, i32 1, i32 0, !dbg !938
  %1 = bitcast i64** %tmp.27.sroa.2.0.cast.1108.sroa_idx13 to i8**, !dbg !938
  store i8* %call.19, i8** %1, align 8, !dbg !938
  %tmp.27.sroa.3.0.cast.1108.sroa_idx14 = getelementptr inbounds { { i8*, i64, i64 }, %IPST.0 }, { { i8*, i64, i64 }, %IPST.0 }* %sret.formal.14, i64 0, i32 1, i32 1, !dbg !938
  store i64 1, i64* %tmp.27.sroa.3.0.cast.1108.sroa_idx14, align 8, !dbg !938
  %tmp.27.sroa.4.0.cast.1108.sroa_idx15 = getelementptr inbounds { { i8*, i64, i64 }, %IPST.0 }, { { i8*, i64, i64 }, %IPST.0 }* %sret.formal.14, i64 0, i32 1, i32 2, !dbg !938
  store i64 1, i64* %tmp.27.sroa.4.0.cast.1108.sroa_idx15, align 8, !dbg !938
  call void @llvm.lifetime.end.p0i8(i64 24, i8* nonnull %"$ret62.0.sroa_cast16"), !dbg !938
  ret void, !dbg !938
}{
entry:
  %"$ret62" = alloca { i8*, i64, i64 }, align 8
  call void @llvm.dbg.value(metadata %TabView.0* %r.19, metadata !933, metadata !DIExpression()), !dbg !934
  %"$ret62.0.sroa_cast16" = bitcast { i8*, i64, i64 }* %"$ret62" to i8*
  call void @llvm.lifetime.start.p0i8(i64 24, i8* nonnull %"$ret62.0.sroa_cast16")
  %"$ret6219" = bitcast { i8*, i64, i64 }* %"$ret62" to i8*
  call void @llvm.memcpy.p0i8.p0i8.i64(i8* nonnull align 8 %"$ret6219", i8* align 16 bitcast ({ i8*, i64, i64 }* @command_line_arguments.fileDescriptor3 to i8*), i64 24, i1 false)
  %call.19 = call i8* @runtime.newobject(i8* nest undef, %_type.0* getelementptr inbounds (%ArrayType.0, %ArrayType.0* @type...61x.7int, i64 0, i32 0)), !dbg !935
  %0 = bitcast i8* %call.19 to i64*, !dbg !935
  store i64 1, i64* %0, align 8, !dbg !935
  call void @llvm.dbg.value(metadata i8* %call.19, metadata !936, metadata !DIExpression(DW_OP_LLVM_fragment, 0, 64)), !dbg !937
  call void @llvm.dbg.value(metadata i64 1, metadata !936, metadata !DIExpression(DW_OP_LLVM_fragment, 64, 64)), !dbg !937
  call void @llvm.dbg.value(metadata i64 1, metadata !936, metadata !DIExpression(DW_OP_LLVM_fragment, 128, 64)), !dbg !937
  %sret.formal.142021 = bitcast { { i8*, i64, i64 }, %IPST.0 }* %sret.formal.14 to i8*
  call void @llvm.memcpy.p0i8.p0i8.i64(i8* nonnull align 8 %sret.formal.142021, i8* nonnull align 8 %"$ret62.0.sroa_cast16", i64 24, i1 false), !dbg !938
  %tmp.27.sroa.2.0.cast.1108.sroa_idx13 = getelementptr inbounds { { i8*, i64, i64 }, %IPST.0 }, { { i8*, i64, i64 }, %IPST.0 }* %sret.formal.14, i64 0, i32 1, i32 0, !dbg !938
  %1 = bitcast i64** %tmp.27.sroa.2.0.cast.1108.sroa_idx13 to i8**, !dbg !938
  store i8* %call.19, i8** %1, align 8, !dbg !938
  %tmp.27.sroa.3.0.cast.1108.sroa_idx14 = getelementptr inbounds { { i8*, i64, i64 }, %IPST.0 }, { { i8*, i64, i64 }, %IPST.0 }* %sret.formal.14, i64 0, i32 1, i32 1, !dbg !938
  store i64 1, i64* %tmp.27.sroa.3.0.cast.1108.sroa_idx14, align 8, !dbg !938
  %tmp.27.sroa.4.0.cast.1108.sroa_idx15 = getelementptr inbounds { { i8*, i64, i64 }, %IPST.0 }, { { i8*, i64, i64 }, %IPST.0 }* %sret.formal.14, i64 0, i32 1, i32 2, !dbg !938
  store i64 1, i64* %tmp.27.sroa.4.0.cast.1108.sroa_idx15, align 8, !dbg !938
  call void @llvm.lifetime.end.p0i8(i64 24, i8* nonnull %"$ret62.0.sroa_cast16"), !dbg !938
  ret void, !dbg !938
}