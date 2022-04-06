package fr.baldir;

import picocli.CommandLine;
import picocli.CommandLine.Command;


@Command(name = "baldir", mixinStandardHelpOptions = true)
public class BaldirCommand implements Runnable {


    @CommandLine.Option(names = "--help", usageHelp = true, description = "display this help and exit")
    boolean help;

    @Override
    public void run() {
        CommandLine.usage(this, System.out);
    }

}
