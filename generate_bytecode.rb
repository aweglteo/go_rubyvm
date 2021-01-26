bytecode = RubyVM::InstructionSequence.compile_file("hello.rb", false).disasm

File.open("bytecode.txt", mode="w"){|f|
 f.write(bytecode)
}
