package fr.baldir;

import io.quarkus.test.junit.main.Launch;
import io.quarkus.test.junit.main.LaunchResult;
import io.quarkus.test.junit.main.QuarkusMainTest;
import org.junit.jupiter.api.Test;

import static org.assertj.core.api.Assertions.assertThat;
@QuarkusMainTest
class Default_command_should {

    @Test
    @Launch("marc")
    void greet_marc_when_marc_is_provided(LaunchResult result){
        assertThat(result.getOutput())
            .contains("Hello marc, go go commando!");
    }

    @Test
    @Launch({})
    void greet_picocli_when_no_name_is_provided(LaunchResult result){
        assertThat(result.getOutput())
            .contains("Hello picocli, go go commando!");
    }
}