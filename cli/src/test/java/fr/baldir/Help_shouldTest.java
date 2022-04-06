package fr.baldir;

import io.quarkus.test.junit.main.Launch;
import io.quarkus.test.junit.main.LaunchResult;
import io.quarkus.test.junit.main.QuarkusMainTest;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import static org.assertj.core.api.Assertions.assertThat;
@QuarkusMainTest
@DisplayName("Help should")
class Help_shouldTest {


    @Test
    @Launch({})
    void be_displayed_when_cli_is_called_with_no_argument(LaunchResult result){
        assertThat(result.getOutput().stripTrailing())
            .isEqualToNormalizingNewlines("Usage: baldir [--help]\n" +
                "      --help   display this help and exit");
    }

    @Test
    @Launch("--help")
    void be_displayed_when_cli_is_called_with_dash_dash_help(LaunchResult result){
        assertThat(result.getOutput().stripTrailing())
            .isEqualToNormalizingNewlines("Usage: baldir [--help]\n" +
                "      --help   display this help and exit");

    }
}