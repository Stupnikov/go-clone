{
entry:
  call void @llvm.dbg.value(metadata %PagerView.0* %m, metadata !280, metadata !DIExpression()), !dbg !281
  %0 = ptrtoint %PagerView.0* %m to i64, !dbg !282
  %call.2 = call { i64, i64 } @github_com_golang_protobuf_proto.CompactTextString(i8* nest undef, i64 ptrtoint ({ %_type.0*, void (i8*, %PagerView.0*)*, void (i8*, %PagerView.0*)*, { i64, i64 } (i8*, %PagerView.0*)* }* @pimt..interface.4ProtoMessage.0func.8.9.8.9.2Reset.0func.8.9.8.9.2String.0func.8.9.8string.9.5..command_line_arguments.PagerView to i64), i64 %0), !dbg !283
  ret { i64, i64 } %call.2, !dbg !284
}{
entry:
  call void @llvm.dbg.value(metadata %PagerView.0* %m, metadata !280, metadata !DIExpression()), !dbg !281
  %0 = ptrtoint %PagerView.0* %m to i64, !dbg !282
  %call.2 = call { i64, i64 } @github_com_golang_protobuf_proto.CompactTextString(i8* nest undef, i64 ptrtoint ({ %_type.0*, void (i8*, %PagerView.0*)*, void (i8*, %PagerView.0*)*, { i64, i64 } (i8*, %PagerView.0*)* }* @pimt..interface.4ProtoMessage.0func.8.9.8.9.2Reset.0func.8.9.8.9.2String.0func.8.9.8string.9.5..command_line_arguments.PagerView to i64), i64 %0), !dbg !283
  ret { i64, i64 } %call.2, !dbg !284
}