{
entry:
  call void @llvm.dbg.value(metadata %ImageView.0* %m, metadata !597, metadata !DIExpression()), !dbg !598
  %0 = ptrtoint %ImageView.0* %m to i64, !dbg !599
  %call.10 = call { i64, i64 } @github_com_golang_protobuf_proto.CompactTextString(i8* nest undef, i64 ptrtoint ({ %_type.0*, void (i8*, %ImageView.0*)*, void (i8*, %ImageView.0*)*, { i64, i64 } (i8*, %ImageView.0*)* }* @pimt..interface.4ProtoMessage.0func.8.9.8.9.2Reset.0func.8.9.8.9.2String.0func.8.9.8string.9.5..command_line_arguments.ImageView to i64), i64 %0), !dbg !600
  ret { i64, i64 } %call.10, !dbg !601
}{
entry:
  call void @llvm.dbg.value(metadata %ImageView.0* %m, metadata !597, metadata !DIExpression()), !dbg !598
  %0 = ptrtoint %ImageView.0* %m to i64, !dbg !599
  %call.10 = call { i64, i64 } @github_com_golang_protobuf_proto.CompactTextString(i8* nest undef, i64 ptrtoint ({ %_type.0*, void (i8*, %ImageView.0*)*, void (i8*, %ImageView.0*)*, { i64, i64 } (i8*, %ImageView.0*)* }* @pimt..interface.4ProtoMessage.0func.8.9.8.9.2Reset.0func.8.9.8.9.2String.0func.8.9.8string.9.5..command_line_arguments.ImageView to i64), i64 %0), !dbg !600
  ret { i64, i64 } %call.10, !dbg !601
}